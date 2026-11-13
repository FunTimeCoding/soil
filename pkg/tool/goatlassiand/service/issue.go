package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Issue(key string) (*issue.Issue, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*issue.Issue, error) { return s.jira.Issue(key) },
	)
}
