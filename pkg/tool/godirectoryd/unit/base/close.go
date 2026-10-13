package base

func (s *Server) Close() {
	s.Authorization.Close()
	s.Stop()
}
