package salt_tester

func (s *Server) Hold() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.hold = make(chan struct{})
	s.holding = make(chan struct{}, 1)
}
