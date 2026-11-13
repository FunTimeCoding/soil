package service

import "context"

func (s *Service) Comment(
	key string,
	body string,
) error {
	return s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error { return s.jira.Comment(key, body) },
	)
}
