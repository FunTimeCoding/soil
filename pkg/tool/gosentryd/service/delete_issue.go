package service

func (s *Service) DeleteIssue(identifier string) error {
	return s.client.DeleteIssue(s.organization, identifier)
}
