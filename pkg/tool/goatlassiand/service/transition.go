package service

import "context"

func (s *Service) Transition(
	key string,
	transition string,
) error {
	return s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error { return s.jira.Transition(key, transition) },
	)
}
