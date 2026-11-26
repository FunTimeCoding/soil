package discussion_server

func (s *Server) Requests() []string {
	s.lock.Lock()
	defer s.lock.Unlock()

	return append([]string{}, s.requests...)
}
