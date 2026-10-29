package base

func (s *Server) Close() {
	s.Service.Close()
	s.Server.Close()
}
