package channel

func (s *Server) Connected() {
	s.readyOnce.Do(func() { close(s.ready) })
}
