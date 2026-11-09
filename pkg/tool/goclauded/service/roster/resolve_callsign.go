package roster

func (s *Service) ResolveCallsign(c string) (string, error) {
	return s.store.ResolveCallsign(c)
}
