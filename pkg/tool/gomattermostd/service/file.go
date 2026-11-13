package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) File(identifier string) ([]byte, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]byte, error) {
			return s.mattermost.File(identifier)
		},
	)
}
