package base

func (s *Server) Close() {
	if s.served != nil {
		s.served.Close()
	}
}
