package service

func (s *Service) DeleteOtherAuthenticationSessions(identifier string) error {
	return s.store.DeleteOtherAuthenticationSessions(identifier)
}
