package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateTag(name string) (*tag.Tag, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tag.Tag, error) {
			return s.client.CreateTag(name)
		},
	)
}
