package roster

func (s *Service) ReleaseCallsign(c string) error {
	if e := s.store.ReleaseCallsign(c); e != nil {
		return e
	}

	s.notify()

	return nil
}
