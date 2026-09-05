package base

func (s *Server) Close() {
	s.store.Close()
	s.server.Stop()
}
