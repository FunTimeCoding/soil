package roster

func (s *Service) ResolveByCallsign(c string) (string, error) {
	return s.store.ResolveByCallsign(c)
}
