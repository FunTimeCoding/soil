package model_context

func (s *Server) register() {
	s.registerOrganization()
	s.registerIssues()
	s.registerEvents()
	s.registerTriage()
}
