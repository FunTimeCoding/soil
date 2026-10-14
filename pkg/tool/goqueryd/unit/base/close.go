package base

func (s *Server) Close() {
	s.store.Close()
	s.embedder.Close()
	s.Server.Close()
}
