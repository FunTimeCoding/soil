package salt_tester

func (s *Server) Logins() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.logins
}
