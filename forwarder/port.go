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
	for {
		var dest_con net.Conn
		dest_con, err = l.Accept()
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
