package fake_assistant

func (s *Server) Dialed() <-chan struct{} {
	return s.dialed
}
