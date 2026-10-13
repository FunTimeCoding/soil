package model_context_server

func (s *Server) Stop() {
	s.pool.Close()
	s.Lifecycle.Stop()
}
