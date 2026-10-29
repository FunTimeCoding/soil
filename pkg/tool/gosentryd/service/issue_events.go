package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) IssueEvents(
	identifier string,
	query string,
	limit int,
	cursor string,
) ([]response.Event, error) {
	return guarded(
		s.guard,
		func() ([]response.Event, error) {
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
