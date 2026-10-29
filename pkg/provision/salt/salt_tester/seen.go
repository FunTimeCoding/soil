package salt_tester

func (s *Server) Seen() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return append([]string{}, s.seen...)
}
