package store

func (s *Store) Close() {
	close(s.stop)
}
