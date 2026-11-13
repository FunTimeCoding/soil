package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreatePlatform(name string) (*platform.Platform, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*platform.Platform, error) {
			return s.client.CreatePlatform(name)
		},
	)
}
