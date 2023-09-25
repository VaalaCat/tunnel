package main

import (
	"context"
	"fmt"
	"net"
	"time"
	"tunnel/client"
	"tunnel/protogen"
	"tunnel/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	go runServer()
	for {
		time.Sleep(1 * time.Second)
		runClient()
	}
}

func runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8989))
	if err != nil {
		panic(err)
	}
	logrus.Info("Server listening on port 8080")

	srv := grpc.NewServer([]grpc.ServerOption{}...)
	protogen.RegisterTunnelServerServer(srv, &server.TunnelServer{})
	srv.Serve(lis)
}

func runClient() {
	cli, conn := client.NewClient()
	defer conn.Close()

	k, err := cli.Connect(context.Background(), &protogen.Package{
		Payload: []byte("hello"),
	})
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(k)
}
