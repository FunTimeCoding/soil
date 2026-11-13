package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) IssueEvents(
	identifier string,
	query string,
	limit int,
	cursor string,
) ([]response.Event, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.Event, error) {
			return s.client.IssueEvents(
				s.organization,
				identifier,
				query,
				limit,
				cursor,
			)
		},
	)
}
