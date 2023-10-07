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

var dest *client.Destination

func main() {
	go runServer()
	time.Sleep(1 * time.Second)

	var err error
	address := fmt.Sprintf("127.0.0.1:%d", config.Get().ClientForwardPort)
	if dest, err = client.NewDestination(address); err != nil {
		panic(err)
	}

	runClient()
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

func runClient() {
	cli, _ := client.NewClient()

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
