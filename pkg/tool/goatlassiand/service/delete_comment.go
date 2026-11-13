package service

import "context"

func (s *Service) DeleteComment(
	key string,
	identifier string,
) error {
	return s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error {
			return s.jira.DeleteComment(key, identifier)
		},
	)
}
