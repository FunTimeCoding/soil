package model_context

func (s *Server) register() {
	s.registerJiraSearch()
	s.registerJiraCreation()
	s.registerJiraWorkflow()
	s.registerJiraComments()
	s.registerJiraLinks()
	s.registerJiraChecklists()
	s.registerConfluenceRead()
	s.registerConfluencePages()
	s.registerConfluenceEdits()
}
