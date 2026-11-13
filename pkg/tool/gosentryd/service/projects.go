package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Projects() ([]response.Project, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.Project, error) {
			return s.client.OrganizationProjects(s.organization)
		},
	)
}
