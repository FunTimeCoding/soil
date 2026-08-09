package service

func (s *Service) AddTags(
	identifier int64,
	tags []string,
) error {
	if e := s.store.AddTags(identifier, tags); e != nil {
		return e
	}

	return s.syncIndexByIdentifier(identifier)
}
