package base

func (s *Server) Close() {
	s.Stop()
}
