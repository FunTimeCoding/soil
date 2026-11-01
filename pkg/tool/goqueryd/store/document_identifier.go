package store

func (s *Store) documentIdentifier(
	collection string,
	path string,
) int {
	var identifier int
	e := s.database.QueryRow(
		"SELECT identifier FROM document WHERE collection = ? AND path = ?",
		collection,
		path,
	).Scan(&identifier)

	if e != nil {
		return 0
	}

	return identifier
}
