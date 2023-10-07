package forwarder

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"tunnel/protogen"

	"github.com/sirupsen/logrus"
)

type Listener interface {
	RegTunnel(*protogen.Tunnel)
	GetTunnel(string) (*net.Listener, error)
}

type ListenerImpl struct {
	ClientMap *sync.Map
}

var globalLis Listener

func (l *ListenerImpl) RegTunnel(tun *protogen.Tunnel) {
	host := "0.0.0.0"
	li, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, tun.GetPort()))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
	l.ClientMap.Store(tun.GetClientId(), &li)
}

func (l *ListenerImpl) GetTunnel(clientID string) (*net.Listener, error) {
	rawli, ok := l.ClientMap.Load(clientID)
	if !ok || rawli == nil {
		return nil, fmt.Errorf("load raw listener faild")
	}

	return rawli.(*net.Listener), nil
}

func GetListener() Listener {
	if globalLis == nil {
		globalLis = &ListenerImpl{
			ClientMap: &sync.Map{},
		}
	}
	return globalLis
}

func ListenAndServe(cli protogen.TunnelServer_CallServer, clientID string) (err error) {
	for {
		lis, err := GetListener().GetTunnel(clientID)
		if err != nil {
			return err
		}
		var dest_con net.Conn
		dest_con, err = (*lis).Accept()
		if err != nil {
			continue
		}
		logrus.Info("new connection")

		go func(dest_con net.Conn) {
			seq := int64(0)
			for {
				recv_data := make([]byte, 1024)
				n, err := dest_con.Read(recv_data)
				if err != nil {
					logrus.Errorf("server recv data from ingress error: %v", err)
					return
				}

				logrus.Infof("server recv data from ingress length: %d", n)

				seq++
				if err := cli.Send(&protogen.Response{
					Seq:     seq,
					Payload: recv_data[:n],
				}); err != nil {
					logrus.Errorf("server send data to cli error: %v", err)
				}
			}
		}(dest_con)

		go func(dest_con net.Conn) {
			for {
				in, err := cli.Recv()
				if err == io.EOF || in.GetSignal() == protogen.Signal_CLOSE {
					dest_con.Close()
					logrus.Infof("server recv data from client should close")
					return
				}

				logrus.Infof("server recv data from client ,length: %d", len(in.Payload))
				if _, err := dest_con.Write(in.Payload); err != nil {
					logrus.Errorf("server send data to ingress error: %v", err)
				} else {
					logrus.Infof("server send data to ingress, length: %d", len(in.Payload))
				}
				logrus.Infof("ingress address : %v", &dest_con)

			}
		}(dest_con)
	}
}
