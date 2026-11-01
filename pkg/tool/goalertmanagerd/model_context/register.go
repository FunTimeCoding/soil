package model_context

func (s *Server) register() {
	s.instances.Register(s.server)
	s.registerAlerts()
	s.registerSilences()
	s.registerStatus()
}
