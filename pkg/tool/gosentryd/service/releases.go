package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Releases(
	query string,
	limit int,
) ([]response.Release, error) {
	return guarded(
		s.guard,
		func() ([]response.Release, error) {
			return s.client.Releases(s.organization, query, limit)
		},
	)
}
