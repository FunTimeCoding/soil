package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Tags() ([]*tag.Tag, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*tag.Tag, error) {
			return s.client.Tags()
		},
	)
}
