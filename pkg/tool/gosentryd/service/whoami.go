package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Whoami() (*response.User, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*response.User, error) {
			return s.client.Whoami()
		},
	)
}
