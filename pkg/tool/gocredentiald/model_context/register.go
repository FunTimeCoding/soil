package model_context

func (s *Server) register() {
	s.registerRead()
	s.registerWrite()
	s.registerEnvironment()
}
