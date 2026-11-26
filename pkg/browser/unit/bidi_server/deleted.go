package bidi_server

func (s *Server) Deleted() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.deleted
}
