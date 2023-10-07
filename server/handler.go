package server

import (
	"context"
	"fmt"
	"tunnel/forwarder"
	"tunnel/protogen"
)

type TunnelServer struct{}

func (t *TunnelServer) Call(srv protogen.TunnelServer_CallServer) error {
	Meta, err := srv.Recv()
	if err != nil || len(Meta.ClientId) == 0 {
		return fmt.Errorf("server get meta error")
	}
	return forwarder.ListenAndServe(srv, Meta.ClientId)
}

func (t *TunnelServer) Register(ctx context.Context, tun *protogen.Tunnel) (*protogen.Tunnel, error) {
	if tun == nil {
		return nil, fmt.Errorf("server get tunnel error")
	}
	forwarder.GetListener().RegTunnel(tun)
	return tun, nil
}
