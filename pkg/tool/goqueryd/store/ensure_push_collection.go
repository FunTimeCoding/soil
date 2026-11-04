package store

func (s *Store) EnsurePushCollection(name string) error {
	_, e := s.database.Exec(
		`INSERT INTO collection (name, path, pattern, type)
		VALUES (?, '', '', 'push')
		ON CONFLICT(name) DO NOTHING`,
		name,
	)

	return e
}
