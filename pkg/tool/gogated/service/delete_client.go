package service

func (s *Service) DeleteClient(identifier string) error {
	return s.store.DeleteClient(identifier)
}
