package base

func (s *Server) Close() {
	s.Store.Close()
	s.Authorization.Close()
	s.Server.Close()
}
