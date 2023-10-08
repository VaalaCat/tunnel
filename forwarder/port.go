package forwarder

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"github.com/VaalaCat/tunnel/protogen"
	"github.com/sirupsen/logrus"
)

type Listener interface {
	RegTunnel(*protogen.Tunnel)
	GetListener(string) (*net.Listener, error)
	GetTunnelInfo(string) (*protogen.Tunnel, error)
	DeleteTunnel(string) error
}

type ListenerImpl struct {
	ClientMap *sync.Map
	InfoMap   *sync.Map
}

var globalLis Listener

func (l *ListenerImpl) RegTunnel(tun *protogen.Tunnel) {
	host := "0.0.0.0"
	li, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, tun.GetPort()))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
	l.ClientMap.Store(tun.GetClientID(), &li)
	l.InfoMap.Store(tun.GetClientID(), tun)
}

func (l *ListenerImpl) GetListener(clientID string) (*net.Listener, error) {
	rawli, ok := l.ClientMap.Load(clientID)
	if !ok || rawli == nil {
		return nil, fmt.Errorf("load raw listener faild")
	}

	return rawli.(*net.Listener), nil
}

func (l *ListenerImpl) GetTunnelInfo(clientID string) (*protogen.Tunnel, error) {
	rawli, ok := l.InfoMap.Load(clientID)
	if !ok || rawli == nil {
		return nil, fmt.Errorf("load raw listener faild")
	}

	return rawli.(*protogen.Tunnel), nil
}

func (l *ListenerImpl) DeleteTunnel(clientID string) error {
	lis, ok := l.ClientMap.Load(clientID)
	if !ok || lis == nil {
		return fmt.Errorf("load raw listener faild")
	}

	if err := (*lis.(*net.Listener)).Close(); err != nil {
		logrus.Errorf("close listener error: %v", err)
	}

	tunnel, ok := l.InfoMap.Load(clientID)
	if !ok || tunnel == nil {
		logrus.Errorf("load tunnel info error")
	}

	l.ClientMap.Delete(clientID)
	l.InfoMap.Delete(clientID)
	logrus.Errorf("delete tunnel: %+v", tunnel)
	return nil
}

func GetListener() Listener {
	if globalLis == nil {
		globalLis = &ListenerImpl{
			ClientMap: &sync.Map{},
			InfoMap:   &sync.Map{},
		}
	}
	return globalLis
}

func ListenAndServe(cli protogen.TunnelServer_CallServer, clientID string) (err error) {
	for {
		lis, err := GetListener().GetListener(clientID)
		if err != nil {
			return err
		}
		var dest_con net.Conn
		dest_con, err = (*lis).Accept()
		if err != nil {
			continue
		}
		tunnel, err := GetListener().GetTunnelInfo(clientID)
		if err != nil {
			return err
		}
		logrus.Infof("server get new connection from client: %+v", tunnel)

		go func(dest_con net.Conn) {
			seq := int64(0)
			for {
				recv_data := make([]byte, 1024)
				n, err := dest_con.Read(recv_data)
				if err != nil {
					logrus.Debugf("server recv data from ingress error: %v", err)
					if err == io.EOF {
						dest_con.Close()
					}
					return
				}

				logrus.Debugf("server recv data from ingress length: %d", n)

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
					logrus.Debugf("server recv data from client should close")
					return
				}

				logrus.Debugf("server recv data from client ,length: %d", len(in.Payload))
				if _, err := dest_con.Write(in.Payload); err != nil {
					logrus.Errorf("server send data to ingress error: %v", err)
				} else {
					logrus.Debugf("server send data to ingress, length: %d", len(in.Payload))
				}
				logrus.Debugf("ingress address : %v", &dest_con)

			}
		}(dest_con)
	}
}
