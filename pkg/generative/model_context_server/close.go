package model_context_server

func (s *Server) Close() {
	s.pool.Close()
	s.Lifecycle.Close()
}
