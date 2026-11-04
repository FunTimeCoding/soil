package store

func (s *Store) SetMetadata(
	collection string,
	path string,
	metadata map[string][]string,
) error {
	var identifier int
	e := s.database.QueryRow(
		"SELECT identifier FROM document WHERE collection = ? AND path = ?",
		collection,
		path,
	).Scan(&identifier)

	if e != nil {
		return nil
	}

	if _, f := s.database.Exec(
		"DELETE FROM document_metadata WHERE document_identifier = ?",
		identifier,
	); f != nil {
		return f
	}

	for key, values := range metadata {
		for _, value := range distinct(values) {
			if _, f := s.database.Exec(
				"INSERT INTO document_metadata (document_identifier, key, value) VALUES (?, ?, ?)",
				identifier,
				key,
				value,
			); f != nil {
				return f
			}
		}
	}

	return nil
}
