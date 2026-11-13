package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (s *Service) View(identifier int) (*view.View, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*view.View, error) {
			return s.client.View(identifier)
		},
	)
}
