package model_context

func (s *Server) register() {
	s.registerWriting()
	s.registerReading()
	s.registerFetching()
	s.registerGraph()
	s.registerTags()
}
