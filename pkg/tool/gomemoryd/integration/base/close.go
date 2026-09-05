package base

func (s *Server) Close() {
	s.server.Stop()
}
