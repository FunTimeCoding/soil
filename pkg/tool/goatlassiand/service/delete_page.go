package service

import "context"

func (s *Service) DeletePage(
	identifier string,
	draft bool,
) error {
	return s.confluenceGuard.Attempt(
		context.Background(),
		func(context.Context) error {
			if draft {
				return s.confluence.DeleteDraft(identifier)
			}

			return s.confluence.Delete(identifier)
		},
	)
}
