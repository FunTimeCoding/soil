package store

func (s *Store) MemoryNamed(
	scope string,
	name string,
) (bool, error) {
	var count int
	e := s.database.QueryRow(
		`SELECT COUNT(*) FROM memory WHERE scope = ? AND name = ? AND is_active = 1`,
		scope,
		name,
	).Scan(&count)

	return count > 0, e
}
