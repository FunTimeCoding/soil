package store

func (s *Store) UniversalTimestamps() bool {
	return timestampsNormalized(s.database)
}
