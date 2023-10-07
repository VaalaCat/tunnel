package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
	"tunnel/client"
	"tunnel/config"
	"tunnel/protogen"
	"tunnel/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	go runServer()
	time.Sleep(1 * time.Second)

	address := fmt.Sprintf("127.0.0.1:%d", config.Get().ClientForwardPort)

	go runClient(int64(7002), address)
	runClient(int64(7005), "127.0.0.1:7004")
}

func runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Get().ServerRPCPort))
	if err != nil {
		panic(err)
	}
	logrus.Infof("Server listening on port %d", config.Get().ServerRPCPort)

	srv := grpc.NewServer([]grpc.ServerOption{}...)
	protogen.RegisterTunnelServerServer(srv, &server.TunnelServer{})
	go srv.Serve(lis)
}

func runClient(port int64, address string) {

	dest, err := client.NewDestination(address)
	if err != nil {
		logrus.Infof("make dest error: %v", err)
		return
	}

	cli, _ := client.NewClient()
	cli.Register(context.Background(), &protogen.Tunnel{
		Port: port,
	})

	c, err := cli.Call(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	go func() {
		for {
			if !dest.CanRead() {
				continue
			}
			recv_data := make([]byte, 2048)
			n, err := dest.Read(recv_data)

			if err == io.EOF {
				logrus.Error("client read data from dest, err:", err)
				dest.Refresh()
				c.Send(&protogen.Request{Signal: protogen.Signal_CLOSE})
				continue
			}

			if err != nil {
				continue
			}

			logrus.Infof("client read data from dest, length: %d", n)
			if err := c.Send(&protogen.Request{Payload: recv_data[:n]}); err != nil {
				logrus.Error("send data to server, err:", err)
			} else {
				logrus.Infof("client send data to server, length: %d", n)
			}
		}
	}()

	for {
		in, err := c.Recv()
		if err == io.EOF {
			logrus.Infof("client get data from server: EOF")
		}
		if err != nil {
			logrus.Errorf("client get data from server: %v", err)
		}

		if err != nil {
			continue
		}
		if in.GetSeq() == 1 {
			dest.Refresh()
		}
		if n, err := dest.Write(in.Payload); err != nil {
			logrus.Error("send in data to d_conn ", err)
		} else {
			logrus.Infof("client send data to dest length: %d", n)
		}
	}
}
