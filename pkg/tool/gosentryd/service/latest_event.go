package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) LatestEvent(issue string) (*response.Event, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*response.Event, error) {
			return s.client.LatestEvent(s.organization, issue)
		},
	)
}
