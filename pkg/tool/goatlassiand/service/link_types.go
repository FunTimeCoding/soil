package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/link_type"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) LinkTypes() ([]link_type.Type, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) ([]link_type.Type, error) {
			return s.jira.LinkTypes()
		},
	)
}
