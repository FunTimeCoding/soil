package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

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

	return s.client.SearchEvents(
		s.organization,
		query,
		identifier,
		limit,
		cursor,
	)
}
