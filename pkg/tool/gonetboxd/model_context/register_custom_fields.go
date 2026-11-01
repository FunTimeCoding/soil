package model_context

func (s *Server) registerCustomFields() {
	s.registerInventory()
	s.registerChoiceSets()
	s.registerFieldValues()
}
