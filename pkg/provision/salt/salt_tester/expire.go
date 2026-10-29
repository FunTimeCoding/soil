package salt_tester

func (s *Server) Expire() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.valid = "expired"
}
