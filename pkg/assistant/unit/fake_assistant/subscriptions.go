package fake_assistant

func (s *Server) Subscriptions() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.subscription
}
