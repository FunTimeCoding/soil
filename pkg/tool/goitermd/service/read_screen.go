package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ReadScreen(identifier string) (*screen.Screen, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*screen.Screen, error) {
			return s.client.ReadScreen(identifier)
		},
	)
}
