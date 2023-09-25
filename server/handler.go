package server

import (
	"context"
	"tunnel/forwarder"
	"tunnel/protogen"
)

type TunnelServer struct{}

func (t *TunnelServer) Call(srv protogen.TunnelServer_CallServer) error {
	return forwarder.ListenAndServe("8080", srv)
}

func (*TunnelServer) Register(ctx context.Context, t *protogen.Tunnel) (*protogen.Tunnel, error) {
	return nil, nil
}
