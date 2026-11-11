package model_context

func (s *Server) register() {
	s.registerModules()
	s.registerReferences()
	s.registerSearch()
	s.registerPatternRewrite()
	s.registerConstructors()
	s.registerNames()
	s.registerRenameSymbols()
	s.registerRenameReceivers()
	s.registerParameters()
	s.registerMoveSymbol()
	s.registerMoveSymbols()
	s.registerMoveMethods()
	s.registerPackages()
	s.registerExtraction()
	s.registerImports()
	s.registerRetarget()
}
