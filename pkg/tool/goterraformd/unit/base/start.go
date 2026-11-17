package base

func (s *Server) Start() {
	s.t.Helper()

	for _, w := range s.Daemon().Workers() {
		w.Start()
	}
}
