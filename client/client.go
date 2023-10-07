package client

import (
	"context"
	"fmt"
	"io"
	"tunnel/protogen"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(rpcPort int64) (protogen.TunnelServerClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", rpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := protogen.NewTunnelServerClient(conn)
	return client, conn
}

func RunClient(rpcPort, ingressPort int64, forwardAddress, clientID string) {

	dest, err := NewDestination(forwardAddress)
	if err != nil {
		logrus.Infof("make dest error: %v", err)
		return
	}

	cli, _ := NewClient(rpcPort)
	cli.Register(context.Background(), &protogen.Tunnel{
		Port:     ingressPort,
		ClientId: clientID,
	})

	c, err := cli.Call(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	c.Send(&protogen.Request{Signal: protogen.Signal_OPEN,
		ClientId: clientID})
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
