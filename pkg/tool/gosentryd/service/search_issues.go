package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SearchIssues(
	query string,
	project string,
	limit int,
	cursor string,
) ([]response.Issue, error) {
	identifier, e := s.resolveProject(project)

	if e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.Issue, error) {
			return s.client.SearchIssues(
				s.organization,
				query,
				identifier,
				limit,
				cursor,
			)
		},
	)
}
