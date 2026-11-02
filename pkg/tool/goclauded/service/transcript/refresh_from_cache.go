package transcript

func (s *Service) RefreshFromCache(identifier string) {
	state, exists := s.cache.Get(identifier)

	if !exists {
		return
	}

	s.RefreshSession(identifier, state)
}
