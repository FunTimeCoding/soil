package model_context

func (s *Server) register() {
	s.instances.Register(s.server)
	s.registerQueries()
	s.registerSeries()
	s.registerRules()
}
