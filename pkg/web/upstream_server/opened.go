package upstream_server

func (s *Server) Opened() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.opened
}
