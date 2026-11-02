package session

func (s *Session) Close() {
	s.login.Close()
}
