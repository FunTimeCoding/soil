package issuer_server

func (s *Server) Discoveries() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.discovery
}
