package base

func (s *Server) Captured() int {
	return len(s.Reporter.Events())
}
