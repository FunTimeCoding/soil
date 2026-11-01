package model_context

func (s *Server) register() {
	s.registerSeeds()
	s.registerDecisions()
	s.registerDecisionFollowUps()
}
