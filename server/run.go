package server

import (
	"fmt"
	"net"
	"tunnel/protogen"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunServer(port int64) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}
	logrus.Infof("Server listening on port %d", port)

	srv := grpc.NewServer([]grpc.ServerOption{}...)
	protogen.RegisterTunnelServerServer(srv, &TunnelServer{})
	go srv.Serve(lis)
}
