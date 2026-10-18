package session

func (s *Session) Token() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.token
}
