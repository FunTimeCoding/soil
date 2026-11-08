package label

func (s *Service) AddLabel(
	sessionIdentifier string,
	name string,
	target string,
	key string,
	value string,
) (string, error) {
	old, e := s.store.AddLabel(sessionIdentifier, key, value)

	if e != nil {
		return "", e
	}

	return s.record(sessionIdentifier, name, target, key, old)
}
