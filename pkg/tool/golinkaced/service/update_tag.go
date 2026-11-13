package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UpdateTag(
	identifier int,
	patch map[string]any,
) (*tag.Tag, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tag.Tag, error) {
			return s.client.UpdateTag(identifier, patch)
		},
	)
}
