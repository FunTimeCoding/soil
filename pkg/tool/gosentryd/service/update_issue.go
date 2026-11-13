package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UpdateIssue(
	identifier string,
	status string,
	assignedTo string,
) (*response.Issue, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*response.Issue, error) {
			return s.client.UpdateIssue(
				s.organization,
				identifier,
				status,
				assignedTo,
			)
		},
	)
}
