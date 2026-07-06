package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) deactivateDocument(
	collection string,
	path string,
) {
	_, e := s.database.Exec(
		"UPDATE document SET active = 0 WHERE collection = ? AND path = ? AND active = 1",
		collection,
		path,
	)
	errors.PanicOnError(e)
}
