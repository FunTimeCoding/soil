package store

func (s *Store) DeleteEmbeddings(hash string) error {
	_, e := s.database.Exec("DELETE FROM embedding WHERE hash = ?", hash)

	return e
}
