package service

func (s *Service) HiddenIdentifiers() (map[int64]bool, error) {
	if s.hiddenTag == "" {
		return map[int64]bool{}, nil
	}

	return s.store.ListIdentifiersByTag(s.hiddenTag)
}
