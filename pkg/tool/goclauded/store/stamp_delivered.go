package store

func (s *Store) StampDelivered() {
	migrateConsumedAt(s.database)
}
