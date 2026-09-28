package service

func (s *Service) DeleteAuthenticationSession(identifier string) error {
	return s.store.DeleteAuthenticationSession(identifier)
}
