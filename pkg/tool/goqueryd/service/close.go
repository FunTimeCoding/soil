package service

func (s *Service) Close() {
	s.guard.Close()
}
