package base

func (s *Server) Close() {
	s.ContextServer.Close()
}
