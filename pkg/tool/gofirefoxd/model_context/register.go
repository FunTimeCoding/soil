package model_context

func (s *Server) register() {
	s.registerTabs()
	s.registerReading()
	s.registerNavigation()
	s.registerGroups()
	s.registerActions()
	s.registerDownloads()
}
