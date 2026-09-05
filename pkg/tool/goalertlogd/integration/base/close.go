package base

func (s *Server) Close() {
	s.Store.Close()
	s.ContextServer.Stop()
}
