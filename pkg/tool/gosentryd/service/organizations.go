package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Organizations() ([]response.Organization, error) {
	return guarded(s.guard, s.client.Organizations)
}
