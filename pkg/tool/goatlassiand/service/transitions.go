package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Transitions(key string) ([]jira.Transition, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) ([]jira.Transition, error) {
			return s.jira.Transitions(key)
		},
	)
}
