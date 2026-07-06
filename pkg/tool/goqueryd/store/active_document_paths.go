package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) activeDocumentPaths(collection string) []string {
	rows, e := s.database.Query(
		"SELECT path FROM document WHERE collection = ? AND active = 1",
		collection,
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var result []string

	for rows.Next() {
		var path string
		errors.PanicOnError(rows.Scan(&path))
		result = append(result, path)
	}

	errors.PanicOnError(rows.Err())

	return result
}
