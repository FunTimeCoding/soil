package base

func (s *Server) Close() {
	s.Web.Stop()
}
