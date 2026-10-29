package service

func (s *Service) Transition(
	key string,
	transition string,
) error {
	return attempt(
		s.jiraGuard,
		func() error { return s.jira.Transition(key, transition) },
	)
}
