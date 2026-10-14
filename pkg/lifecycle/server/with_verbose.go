package server

func (s *Server) WithVerbose(v bool) *Server {
	s.verbose = v

	return s
}
