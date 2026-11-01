package fake_assistant

func (s *Server) RefuseToken() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.invalid = true
}
