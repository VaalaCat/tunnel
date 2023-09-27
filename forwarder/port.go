package forwarder

import (
	"fmt"
	"io"
	"net"
	"os"
	"tunnel/config"
	"tunnel/protogen"

	"github.com/sirupsen/logrus"
)

var dest_con net.Conn
var l net.Listener

func init() {
	var err error
	host := "0.0.0.0"
	l, err = net.Listen("tcp", fmt.Sprintf("%s:%d", host, config.Get().ServerIngressPort))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
}

func ListenAndServe(cli protogen.TunnelServer_CallServer) (err error) {
	done := make(chan bool)
	for {
		select {
		case <-done:
			return nil
		default:
			dest_con, err = l.Accept()
			if err != nil {
				continue
			}
			logrus.Info("new connection")

			go func(dest_con net.Conn) {
				for {
					recv_data := make([]byte, 1024)
					n, err := dest_con.Read(recv_data)
					if err == io.EOF {
						logrus.Errorf("server recv data from ingress error: %v", err)
						done <- true
						return
					}
					if err != nil {
						logrus.Errorf("server recv data from ingress error: %v", err)
						done <- true
						return
					}

					logrus.Infof("server recv data from ingress length: %d", n)

					if err := cli.Send(&protogen.Response{Payload: recv_data[:n]}); err != nil {
						logrus.Errorf("server send data to cli error: %v", err)
						continue
					}
				}
			}(dest_con)
			go func(dest_con net.Conn) {
				for {
					in, err := cli.Recv()
					if err == io.EOF {
						logrus.Errorf("server recv data from client error: %v", err)
						done <- true
						return
					}
					if err != nil {
						logrus.Errorf("server recv data from client error: %v", err)
						done <- true
						return
					}
					logrus.Infof("server recv data from client ,length: %d", len(in.Payload))
					if _, err := dest_con.Write(in.Payload); err != nil {
						logrus.Errorf("server send data to ingress error: %v", err)
						continue
					}
				}
			}(dest_con)
		}
	}
}
