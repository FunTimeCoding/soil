package service

func (s *Service) DeletePage(
	identifier string,
	draft bool,
) error {
	return attempt(
		s.confluenceGuard,
		func() error {
			if draft {
				return s.confluence.DeleteDraft(identifier)
			}

			return s.confluence.Delete(identifier)
		},
	)
}
