package roster

func (s *Service) SessionCallsign(identifier string) (string, error) {
	return s.store.SessionCallsign(identifier)
}
