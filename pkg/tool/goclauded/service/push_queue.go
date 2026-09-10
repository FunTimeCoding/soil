package service

func (s *Service) PushQueue(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) error {
	if e := s.store.PushQueue(
		sessionIdentifier,
		callsign,
		kind,
		body,
	); e != nil {
		return e
	}

	s.notify()

	return nil
}
