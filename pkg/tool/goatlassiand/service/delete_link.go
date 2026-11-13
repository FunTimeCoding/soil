package service

import "context"

func (s *Service) DeleteLink(identifier string) error {
	return s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error { return s.jira.DeleteLink(identifier) },
	)
}
