package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Event(
	project string,
	identifier string,
) (*response.Event, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*response.Event, error) {
			return s.client.Event(s.organization, project, identifier)
		},
	)
}
