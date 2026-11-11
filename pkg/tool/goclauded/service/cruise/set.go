package cruise

func (s *Service) Set(
	sessionIdentifier string,
	text string,
) error {
	grace, e := ParseGrace(text)

	if e != nil {
		return e
	}

	if f := s.store.SetCruise(sessionIdentifier, grace); f != nil {
		return f
	}

	s.notifier.Notify()

	return nil
}
