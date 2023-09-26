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
	// Refresh()
	n := 0
	for {
		n++
		logrus.Infof("start client count: %d", n)
		runClient()
	}
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
	// defer c.CloseSend()

	go func() {
		for {
			recv_data := make([]byte, 2048)
			n, err := GetConnection().Read(recv_data)
			if err == io.EOF {
				logrus.Error("client read data from dest, err:", err)
				break
			}

			if err != nil {
				logrus.Error("client read data from dest, err:", err)
				break
			}

			logrus.Infof("client read data from dest, length: %d", len(recv_data))
			if err := c.Send(&protogen.Request{Payload: recv_data[:n]}); err != nil {
				logrus.Error("send data to server, err:", err)
				continue
			} else {
				logrus.Infof("client send data to server, length: %d", len(recv_data))
			}
		}
	}()

	for {
		in, err := c.Recv()
		if err == io.EOF {
			logrus.Infof("client get data from server: EOF")
			break
		}
		if err != nil {
			logrus.Errorf("client get data from server: %v", err)
			break
		}
		logrus.Infof("client get data from server, length: %d", len(in.Payload))
		if n, err := GetConnection().Write(in.Payload); err != nil {
			logrus.Error("send in data to d_conn ", err)
			continue
		} else {
			logrus.Infof("client send data to dest length: %d", n)
		}
	}
}

var conn *net.TCPConn

func GetConnection() *net.TCPConn {
	if conn != nil {
		return conn
	}
	d_tcpAddr, err := net.ResolveTCPAddr("tcp4",
		fmt.Sprintf("127.0.0.1:%d", config.Get().ClientForwardPort))
	if err != nil {
		logrus.Errorf("resolv tcp error :%v", err)
		return nil
	}

	d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
	if err != nil {
		logrus.Errorf("dial tcp error: %v", err)
		return nil
	}
	d_conn.SetKeepAlive(true)
	d_conn.SetKeepAlivePeriod(30 * time.Second)
	d_conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	d_conn.SetWriteBuffer(2048)

	conn = d_conn
	return conn
}

func Refresh() {
	go func() {
		for {
			d_tcpAddr, err := net.ResolveTCPAddr("tcp4",
				fmt.Sprintf("127.0.0.1:%d", config.Get().ClientForwardPort))
			if err != nil {
				logrus.Errorf("resolv tcp error :%v", err)
				return
			}

			d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
			if err != nil {
				logrus.Errorf("dial tcp error: %v", err)
				return
			}
			d_conn.SetKeepAlive(true)
			d_conn.SetKeepAlivePeriod(30 * time.Second)
			d_conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			d_conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			d_conn.SetWriteBuffer(2048)

			conn = d_conn
			time.Sleep(10 * time.Second)
		}
	}()
}
