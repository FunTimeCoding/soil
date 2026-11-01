package model_context

func (s *Server) register() {
	s.registerDevices()
	s.registerInterfaces()
	s.registerAddressing()
	s.registerCatalog()
	s.registerClusters()
	s.registerVirtualMachines()
	s.registerTags()
	s.registerJournal()
	s.registerJournalEdits()
	s.registerLabels()
	s.registerTunnels()
	s.registerTunnelTerminations()
	s.registerConfigurationContexts()
	s.registerCustomFields()
}
