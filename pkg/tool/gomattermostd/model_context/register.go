package model_context

func (s *Server) register() {
	s.registerUser()
	s.registerSearch()
	s.registerFiles()
	s.registerChannels()
	s.registerMessages()
	s.registerDirectory()

	if s.service.Subscribing() {
		s.registerSubscription()
	}

	if s.service.Monitoring() {
		s.registerMonitoring()
	}
}
