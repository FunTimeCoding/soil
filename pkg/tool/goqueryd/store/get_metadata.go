package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) GetMetadata(documentIdentifier int) map[string]string {
	rows, e := s.database.Query(
		"SELECT key, value FROM document_metadata WHERE document_identifier = ?",
		documentIdentifier,
	)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(rows)
	result := map[string]string{}

	for rows.Next() {
		var key, value string
		errors.PanicOnError(rows.Scan(&key, &value))
		result[key] = value
	}

	errors.PanicOnError(rows.Err())

	if len(result) == 0 {
		return nil
	}

	return result
}
