package store

func (s *Store) NormalizeTimestamps() {
	migrateTimestamps(s.database)
}
