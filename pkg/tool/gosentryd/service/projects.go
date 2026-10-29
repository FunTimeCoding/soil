package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Projects() ([]response.Project, error) {
	return guarded(
		s.guard,
		func() ([]response.Project, error) {
			return s.client.OrganizationProjects(s.organization)
		},
	)
}
