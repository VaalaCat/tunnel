package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"tunnel/client"
	"tunnel/protogen"
	"tunnel/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	go runServer()
	runClient()
}

func runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8989))
	if err != nil {
		panic(err)
	}
	logrus.Info("Server listening on port 8080")

	srv := grpc.NewServer([]grpc.ServerOption{}...)
	protogen.RegisterTunnelServerServer(srv, &server.TunnelServer{})
	go srv.Serve(lis)
}

func runClient() {
	cli, conn := client.NewClient()
	defer conn.Close()

	c, err := cli.Call(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	for {
		in, err := c.Recv()
		if err == io.EOF {
			logrus.Info("EOF")
			return
		}
		if err != nil {
			logrus.Error(err)
			return
		}

		d_tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8899")
		if err != nil {
			logrus.Error(err)
			return
		}

		d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
		if err != nil {
			logrus.Error(err)
			return
		}

		go func(d_conn *net.TCPConn) {
			defer d_conn.Close()

			for {
				recv_data := make([]byte, 1024)
				n, err := d_conn.Read(recv_data)
				if err != nil {
					if err != io.EOF {
						logrus.Error("read data from d_conn", err)
					}
					break
				}

				if err := c.Send(&protogen.Request{Payload: recv_data[:n]}); err != nil {
					logrus.Error("send data to server", err)
					break
				}
			}
		}(d_conn)

		if _, err := d_conn.Write(in.Payload); err != nil {
			logrus.Error("send in data to d_conn", err)
			d_conn.Close()
			return
		}
	}
}
