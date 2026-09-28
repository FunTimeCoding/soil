package service

func (s *Service) DeleteLoginSession(identifier string) error {
	return s.store.DeleteLoginSession(identifier)
}
