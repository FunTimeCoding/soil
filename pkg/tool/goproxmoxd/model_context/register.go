package model_context

func (s *Server) register() {
	s.instances.Register(s.server)
	s.registerNodes()
	s.registerStorage()
	s.registerMachines()
	s.registerMachinePower()
	s.registerMachineSnapshots()
	s.registerMachineConfiguration()
	s.registerMachineCopies()
	s.registerContainers()
	s.registerContainerSnapshots()
	s.registerSnippets()
}
