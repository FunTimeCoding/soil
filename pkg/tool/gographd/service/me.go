package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) me() (*user.User, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*user.User, error) {
			return s.client.Me()
		},
	)
}
