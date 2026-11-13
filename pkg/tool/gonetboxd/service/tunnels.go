package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Tunnels() ([]*tunnel.Tunnel, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*tunnel.Tunnel, error) {
			return s.client.Tunnels()
		},
	)
}
