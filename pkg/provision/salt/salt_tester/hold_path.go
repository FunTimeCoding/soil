package salt_tester

func (s *Server) holdPath(path string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.held = path
	s.hold = make(chan struct{})
	s.holding = make(chan struct{}, 1)
}
