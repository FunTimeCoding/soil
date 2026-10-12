package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Organizations() ([]response.Organization, error) {
	return s.client.Organizations()
}
