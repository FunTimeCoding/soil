package fake_assistant

func (s *Server) Refuse(upgrades int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refusal = upgrades
}
