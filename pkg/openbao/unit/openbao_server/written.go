package openbao_server

func (s *Server) Written() map[string]any {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.written
}
