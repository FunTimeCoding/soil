package model_context

func (s *Server) register() {
	s.registerSearch()
	s.registerIndexing()
	s.registerCollections()
	s.registerContexts()
	s.registerDocuments()
	s.registerTags()
	s.registerMetadata()
}
