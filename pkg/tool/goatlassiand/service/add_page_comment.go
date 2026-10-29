package service

func (s *Service) AddPageComment(
	identifier string,
	body string,
) error {
	return attempt(
		s.confluenceGuard,
		func() error { return s.confluence.AddComment(identifier, body) },
	)
}
