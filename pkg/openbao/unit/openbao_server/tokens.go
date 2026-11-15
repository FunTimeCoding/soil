package openbao_server

func (s *Server) Tokens() []string {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.tokens
}
