package service

func (s *Service) RemoveTags(
	identifier int64,
	tags []string,
) error {
	if e := s.store.RemoveTags(identifier, tags); e != nil {
		return e
	}

	return s.syncIndexByIdentifier(identifier)
}
