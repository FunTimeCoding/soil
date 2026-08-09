package service

func (s *Service) ReplaceTags(
	identifier int64,
	tags []string,
) error {
	if e := s.store.ReplaceTags(identifier, tags); e != nil {
		return e
	}

	return s.syncIndexByIdentifier(identifier)
}
