package salt_tester

func (s *Server) Holding() <-chan struct{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.holding
}
