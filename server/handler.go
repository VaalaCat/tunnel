package server

import (
	"context"
	"fmt"

	"github.com/VaalaCat/tunnel/forwarder"
	"github.com/VaalaCat/tunnel/protogen"
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

func (*TunnelServer) Delete(ctx context.Context, req *protogen.DeleteRequest) (*protogen.DeleteResponse, error) {
	err := forwarder.GetListener().DeleteTunnel(req.GetClientID())
	return &protogen.DeleteResponse{Success: err == nil}, err
}

func (*TunnelServer) QueryTunnel(ctx context.Context, req *protogen.QueryTunnelRequest) (*protogen.QueryTunnelResponse, error) {
	tun, err := forwarder.GetListener().GetTunnelInfo(req.GetClientID())
	if err != nil || tun == nil {
		return nil, fmt.Errorf("query tunnel error")
	}
	return &protogen.QueryTunnelResponse{Tunnel: tun}, nil
}
