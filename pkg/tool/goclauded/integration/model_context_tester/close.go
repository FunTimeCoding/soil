package model_context_tester

func (s *Session) Close() {
	s.Client.Close()
}
