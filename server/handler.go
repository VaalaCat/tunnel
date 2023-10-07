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
	if err != nil || len(Meta.ClientID) == 0 {
		return fmt.Errorf("server get meta error")
	}
	return forwarder.ListenAndServe(srv, Meta.ClientID)
}

func (t *TunnelServer) Register(ctx context.Context, tun *protogen.Tunnel) (*protogen.Tunnel, error) {
	if tun == nil {
		return nil, fmt.Errorf("server get tunnel error")
	}
	forwarder.GetListener().RegTunnel(tun)
	return tun, nil
}

func (*TunnelServer) Delete(context.Context, *protogen.DeleteRequest) (*protogen.DeleteResponse, error) {
	return nil, nil
}
