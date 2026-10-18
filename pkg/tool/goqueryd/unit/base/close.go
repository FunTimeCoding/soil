package base

func (s *Server) Close() {
	s.service.Close()
	s.store.Close()
	s.embedder.Close()
	s.Server.Close()
}
