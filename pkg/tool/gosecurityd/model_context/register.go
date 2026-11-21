package model_context

func (s *Server) register() {
	s.registerResources()
	s.registerAdvisory()
	s.registerFlags()
	s.registerSystems()
	s.registerUpgrades()
	s.registerRemediation()
}
