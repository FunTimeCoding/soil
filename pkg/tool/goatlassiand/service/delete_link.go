package service

func (s *Service) DeleteLink(identifier string) error {
	return attempt(
		s.jiraGuard,
		func() error { return s.jira.DeleteLink(identifier) },
	)
}
