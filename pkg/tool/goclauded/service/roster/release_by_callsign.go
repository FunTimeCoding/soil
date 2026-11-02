package roster

func (s *Service) ReleaseByCallsign(c string) error {
	if e := s.store.ReleaseByCallsign(c); e != nil {
		return e
	}

	s.notify()

	return nil
}
