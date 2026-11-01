package model_context

func (s *Server) register() {
	s.instances.Register(s.server)
	s.registerProjects()
	s.registerRepository()
	s.registerPipelines()
	s.registerJobs()
	s.registerMergeRequests()
	s.registerCommits()
	s.registerVariables()
	s.registerVariableChanges()
	s.registerRegistry()
}
