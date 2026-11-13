package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Projects() (*jira.ProjectList, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*jira.ProjectList, error) {
			return s.jira.Projects()
		},
	)
}
