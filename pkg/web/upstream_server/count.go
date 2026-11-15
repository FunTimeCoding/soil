package upstream_server

func (s *Server) Count() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.requests)
}
