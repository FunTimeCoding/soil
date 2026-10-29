package base

func (s *Server) Close() {
	s.service.Close()
	s.Server.Close()
}
