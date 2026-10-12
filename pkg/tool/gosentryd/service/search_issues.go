package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

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

	return s.client.SearchIssues(
		s.organization,
		query,
		identifier,
		limit,
		cursor,
	)
}
