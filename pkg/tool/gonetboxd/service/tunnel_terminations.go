package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) TunnelTerminations() ([]*tunnel_termination.Termination, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*tunnel_termination.Termination, error) {
			return s.client.TunnelTerminations()
		},
	)
}
