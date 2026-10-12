package service

func (s *Service) SetActiveInstance(
	sessionIdentifier string,
	instance string,
) {
	s.sessions.Store(sessionIdentifier, instance)
}
