package session

func (s *Session) Close() {
	s.client.Close()
}
