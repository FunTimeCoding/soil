package service

func (s *Service) Comment(
	key string,
	body string,
) error {
	return attempt(
		s.jiraGuard,
		func() error { return s.jira.Comment(key, body) },
	)
}
