package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/host"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Hosts(phrase string) ([]*host.Host, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*host.Host, error) {
			return s.client.Hosts(phrase)
		},
	)
}
