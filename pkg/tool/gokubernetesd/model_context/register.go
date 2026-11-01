package model_context

func (s *Server) register() {
	s.registerClusters()
	s.registerObjects()
	s.registerLogs()
	s.registerEvents()
	s.registerHealth()
}
