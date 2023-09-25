package client

import (
	"tunnel/protogen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (protogen.TunnelServerClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("127.0.0.1:8989",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := protogen.NewTunnelServerClient(conn)
	return client, conn
}
