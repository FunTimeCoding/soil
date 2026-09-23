package service

func (s *Service) StampTurnEnd(sessionIdentifier string) error {
	return s.store.StampTurnEnd(sessionIdentifier, s.clock())
}
