package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ConfigurationContexts() ([]*configuration_context.Context, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*configuration_context.Context, error) {
			return s.client.ConfigurationContexts()
		},
	)
}
