package service

import "context"

func (s *Service) DeleteIssue(identifier string) error {
	return s.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return s.client.DeleteIssue(s.organization, identifier)
		},
	)
}
