package score

func (s *Service) Generation() uint64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.generation
}
