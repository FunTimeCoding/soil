package salt_tester

func (s *Server) Release() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	close(s.hold)
}
