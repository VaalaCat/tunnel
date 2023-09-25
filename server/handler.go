package server

import (
	"io"
	"tunnel/protogen"

	"github.com/sirupsen/logrus"
)

type TunnelServer struct{}

func (t *TunnelServer) Call(srv protogen.TunnelServer_CallServer) error {
	for {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logrus.Info(string(in.Payload))
		if err := srv.Send(&protogen.Response{Payload: in.Payload}); err != nil {
			return err
		}
	}
}
