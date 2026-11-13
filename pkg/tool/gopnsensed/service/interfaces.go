package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Interfaces() ([]*network_interface.Interface, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*network_interface.Interface, error) {
			return s.client.Interfaces()
		},
	)
}
