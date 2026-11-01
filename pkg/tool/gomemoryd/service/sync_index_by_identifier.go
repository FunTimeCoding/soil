package service

func (s *Service) syncIndexByIdentifier(identifier int64) error {
	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return e
	}

	s.index(m)

	return nil
}
