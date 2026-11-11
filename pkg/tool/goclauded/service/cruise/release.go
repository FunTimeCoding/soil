package cruise

func (s *Service) Release() []error {
	held, e := s.store.HeldCruiseSessions()

	if e != nil {
		return []error{e}
	}

	var result []error
	released := false

	for _, identifier := range held {
		promoted, f := s.releaseHeld(identifier)

		if f != nil {
			result = append(result, f)
		}

		released = released || promoted
	}

	if released {
		s.notifier.Notify()
	}

	return result
}
