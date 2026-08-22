package service

func (s *Service) DeleteLink(identifier string) error {
	return s.jira.DeleteLink(identifier)
}
