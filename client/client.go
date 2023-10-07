package client

import (
	"context"
	"fmt"
	"io"

	"github.com/VaalaCat/tunnel/protogen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(rpcHost string, rpcPort int64) (protogen.TunnelServerClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", rpcHost, rpcPort),
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
		logrus.Debugf("make dest error: %v", err)
		return
	}

	cli, _ := NewClient(rpcPort)
	cli.Register(context.Background(), &protogen.Tunnel{
		Port:     ingressPort,
		ClientID: clientID,
	})

	c, err := cli.Call(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	c.Send(&protogen.Request{Signal: protogen.Signal_OPEN,
		ClientID: clientID})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				logrus.Warnf("client:%v get context done, err: %v", clientID, ctx.Err())
				return
			default:
				recv_data := make([]byte, 2048)
				n, err := dest.Read(recv_data)

				if err == io.EOF {
					logrus.Debugf("client read data from dest, err: %v", err)
					dest.Refresh()
					c.Send(&protogen.Request{Signal: protogen.Signal_CLOSE})
					continue
				}

				if err != nil {
					logrus.Debugf("client read data from dest, err: %v", err)
					dest, err = NewDestination(forwardAddress)
					if err != nil {
						logrus.Debugf("make dest error: %v", err)
						return
					}
					continue
				}

				logrus.Debugf("client read data from dest, length: %d", n)
				if err := c.Send(&protogen.Request{Payload: recv_data[:n]}); err != nil {
					logrus.Error("send data to server, err:", err)
				} else {
					logrus.Debugf("client send data to server, length: %d", n)
				}
			}
		}
	}()

	for {
		in, err := c.Recv()
		if err == io.EOF {
			logrus.Debugf("client get data from server: EOF")
		}
		if err != nil {
			logrus.Debugf("client get data from server: %v", err)
			cancel()
			return
		}

		if err != nil {
			continue
		}
		if n, err := dest.Write(in.Payload); err != nil {
			logrus.Error("client send data to dest error: ", err)
		} else {
			logrus.Debugf("client send data to dest length: %d", n)
		}
	}
}

func DeleteClient(rpcPort int64, clientID string) {
	cli, _ := NewClient(rpcPort)
	cli.Delete(context.Background(), &protogen.DeleteRequest{ClientID: clientID})
}
