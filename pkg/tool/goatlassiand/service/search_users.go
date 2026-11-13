package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SearchUsers(query string) ([]jira.User, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) ([]jira.User, error) {
			return s.jira.FindUsers(query)
		},
	)
}
