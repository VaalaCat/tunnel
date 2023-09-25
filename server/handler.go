package server

import (
	"context"
	"tunnel/protogen"

	"github.com/sirupsen/logrus"
)

type TunnelServer struct{}

func (t *TunnelServer) Connect(ctx context.Context, pkg *protogen.Package) (*protogen.Package, error) {
	logrus.Info("Connect")
	return pkg, nil
}

func (t *TunnelServer) Disconnect(ctx context.Context, pkg *protogen.Package) (*protogen.Package, error) {
	logrus.Info("Disconnect")
	return pkg, nil
}

func (t *TunnelServer) Data(srv protogen.TunnelServer_DataServer) error {
	logrus.Info("Data")
	return nil
}
