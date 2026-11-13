package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Project(key string) (*jira.Project, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*jira.Project, error) {
			return s.jira.Project(key)
		},
	)
}
