package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) collectionPath(name string) (string, string) {
	row := s.database.QueryRow(
		"SELECT path, pattern FROM collection WHERE name = ?",
		name,
	)
	var path, pattern string
	e := row.Scan(&path, &pattern)
	errors.PanicOnError(e)

	return path, pattern
}
