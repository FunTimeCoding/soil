package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Platforms() ([]*platform.Platform, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*platform.Platform, error) {
			return s.client.Platforms()
		},
	)
}
