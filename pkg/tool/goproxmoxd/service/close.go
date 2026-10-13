package service

func (s *Service) Close() {
	s.pool.Close()
	s.snippets.Close()
}
