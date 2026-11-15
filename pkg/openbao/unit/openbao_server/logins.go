package openbao_server

func (s *Server) Logins() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.logins
}
