package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) removeDocument(
	collection string,
	path string,
) {
	_, e := s.database.Exec(
		"DELETE FROM document WHERE collection = ? AND path = ?",
		collection,
		path,
	)
	errors.PanicOnError(e)
}
