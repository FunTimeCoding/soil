package service

func (s *Service) Close() {
	s.runner.Close()
}
