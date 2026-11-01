package model_context

func (s *Server) register() {
	s.registerLinks()
	s.registerListsAndTags()
	s.registerNotes()
	s.registerAcrossTypes()
}
