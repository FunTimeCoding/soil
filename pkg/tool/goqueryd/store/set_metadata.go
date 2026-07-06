package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) SetMetadata(
	collection string,
	path string,
	metadata map[string]string,
) {
	var identifier int
	e := s.database.QueryRow(
		"SELECT identifier FROM document WHERE collection = ? AND path = ? AND active = 1",
		collection,
		path,
	).Scan(&identifier)

	if e != nil {
		return
	}

	_, e = s.database.Exec(
		"DELETE FROM document_metadata WHERE document_identifier = ?",
		identifier,
	)
	errors.PanicOnError(e)

	for key, value := range metadata {
		_, e = s.database.Exec(
			"INSERT INTO document_metadata (document_identifier, key, value) VALUES (?, ?, ?)",
			identifier,
			key,
			value,
		)
		errors.PanicOnError(e)
	}
}
