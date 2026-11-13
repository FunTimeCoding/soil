package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Releases(
	query string,
	limit int,
) ([]response.Release, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.Release, error) {
			return s.client.Releases(s.organization, query, limit)
		},
	)
}
