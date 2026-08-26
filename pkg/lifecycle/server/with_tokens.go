package server

func (s *Server) WithTokens(tokens []string) *Server {
	s.tokens = tokens

	return s
}
