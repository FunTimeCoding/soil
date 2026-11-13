package service

import "context"

func (s *Service) AddPageComment(
	identifier string,
	body string,
) error {
	return s.confluenceGuard.Attempt(
		context.Background(),
		func(context.Context) error {
			return s.confluence.AddComment(identifier, body)
		},
	)
}
