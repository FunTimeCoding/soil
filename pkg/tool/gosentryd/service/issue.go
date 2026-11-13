package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Issue(identifier string) (*response.Issue, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*response.Issue, error) {
			return s.client.IssueByIdentifier(s.organization, identifier)
		},
	)
}
