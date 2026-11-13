package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) updateLink(
	identifier int,
	patch map[string]any,
) (*link.Link, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*link.Link, error) {
			return s.client.UpdateLink(identifier, patch)
		},
	)
}
