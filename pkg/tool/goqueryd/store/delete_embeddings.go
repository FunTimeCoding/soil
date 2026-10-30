package store

func (s *Store) DeleteEmbeddings(hash string) error {
	if _, e := s.database.Exec(
		"DELETE FROM embedding WHERE hash = ?",
		hash,
	); e != nil {
		return e
	}

	_, e := s.database.Exec(
		"DELETE FROM embedding_failure WHERE hash = ?",
		hash,
	)

	return e
}
