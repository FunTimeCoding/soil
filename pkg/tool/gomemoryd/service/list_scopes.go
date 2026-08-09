package service

func (s *Service) ListScopes() ([]string, error) {
	return s.store.ListScopes()
}
