package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateTunnel(
	name string,
	encapsulation string,
	group string,
) (*tunnel.Tunnel, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tunnel.Tunnel, error) {
			g, e := s.client.TunnelGroup(group)

			if e != nil {
				return nil, e
			}

			return s.client.CreateTunnel(name, encapsulation, g)
		},
	)
}
