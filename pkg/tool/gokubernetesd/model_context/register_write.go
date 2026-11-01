package model_context

func (s *Server) registerWrite() {
	s.registerExec()
	s.registerChanges()
	s.registerPortForwards()
}
