package fake_assistant

func (s *Server) Subscribed() <-chan struct{} {
	return s.subscribed
}
