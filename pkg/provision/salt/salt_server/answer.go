package salt_server

func (s *Server) Answer(
	minion string,
	output string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.answers[minion] = output
}
