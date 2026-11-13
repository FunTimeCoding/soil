package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SearchEvents(
	query string,
	project string,
	limit int,
	cursor string,
) ([]response.EventRow, error) {
	identifier, e := s.resolveProject(project)

	if e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.EventRow, error) {
			return s.client.SearchEvents(
				s.organization,
				query,
				identifier,
				limit,
				cursor,
			)
		},
	)
}
