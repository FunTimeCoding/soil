package label

func (s *Service) GetLabel(
	sessionIdentifier string,
	key string,
) (string, error) {
	return s.store.GetLabel(sessionIdentifier, key)
}
