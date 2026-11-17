package model_context_server

func (s *Server) Close() {
	s.closed.Do(
		func() {
			s.pool.Close()
			s.Lifecycle.Close()
		},
	)
}
