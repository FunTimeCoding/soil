package service

func (s *Service) DeleteComment(
	key string,
	identifier string,
) error {
	return s.jira.DeleteComment(key, identifier)
}
