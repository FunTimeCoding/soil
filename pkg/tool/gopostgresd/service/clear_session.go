package service

func (s *Service) ClearSession(sessionIdentifier string) {
	s.sessions.Delete(sessionIdentifier)
}
