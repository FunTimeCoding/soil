package base

func (s *Server) Started() *Server {
	s.t.Helper()

	for _, w := range s.Daemon().Workers() {
		w.Start()
	}

	return s
}
