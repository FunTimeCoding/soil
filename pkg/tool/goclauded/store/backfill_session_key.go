package store

func (s *Store) BackfillSessionKey() {
	migrateSessionKey(s.database)
}
