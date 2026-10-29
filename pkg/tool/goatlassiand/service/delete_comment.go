package service

func (s *Service) DeleteComment(
	key string,
	identifier string,
) error {
	return attempt(
		s.jiraGuard,
		func() error { return s.jira.DeleteComment(key, identifier) },
	)
}
