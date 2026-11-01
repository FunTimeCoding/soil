package model_context

func (s *Server) register() {
	s.registerTabs()
	s.registerReading()
	s.registerNavigation()
	s.registerActions()
}
