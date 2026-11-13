package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UpdateList(
	identifier int,
	patch map[string]any,
) (*list.List, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*list.List, error) {
			return s.client.UpdateList(identifier, patch)
		},
	)
}
