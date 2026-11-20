package model_context

func (s *Server) register() {
	s.registerAdvisory()
	s.registerFlags()
	s.registerSystems()
}
