package message

func (s *Service) SetCruise(
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

	s.notify()

	return nil
}
