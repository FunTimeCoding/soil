package salt_server

func (s *Server) Logins() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.logins
}
