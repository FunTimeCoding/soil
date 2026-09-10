package store

func (s *Store) CountPoolNames() (int, error) {
	names, e := s.poolNames()

	if e != nil {
		return 0, e
	}

	return len(names), nil
}
