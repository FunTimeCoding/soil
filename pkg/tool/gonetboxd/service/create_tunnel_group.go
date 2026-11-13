package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateTunnelGroup(name string) (*tunnel_group.Group, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tunnel_group.Group, error) {
			return s.client.CreateTunnelGroup(name)
		},
	)
}
