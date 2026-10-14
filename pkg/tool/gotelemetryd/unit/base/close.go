package base

func (s *Server) Close() {
	s.Server.Close()
	s.Store.Close()
}
