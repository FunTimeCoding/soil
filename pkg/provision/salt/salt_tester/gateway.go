package salt_tester

func (s *Server) Gateway(status int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.gateway = status
}
