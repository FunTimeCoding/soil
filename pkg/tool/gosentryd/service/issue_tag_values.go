package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) IssueTagValues(
	identifier string,
	tag string,
	limit int,
) ([]response.TagValue, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.TagValue, error) {
			return s.client.IssueTagValues(
				s.organization,
				identifier,
				tag,
				limit,
			)
		},
	)
}
