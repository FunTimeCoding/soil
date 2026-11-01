package model_context

func (s *Server) register() {
	s.registerAuthorities()
	s.registerIssuance()
	s.registerCertificates()
	s.registerPublication()
	s.registerExport()
}
