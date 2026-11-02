package service

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.guard.Close()
			s.memory.Close()
		},
	)
}
