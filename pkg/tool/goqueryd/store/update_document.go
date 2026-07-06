package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) updateDocument(
	identifier int,
	title string,
	hash string,
	now string,
) {
	_, e := s.database.Exec(
		"UPDATE document SET title = ?, hash = ?, modified_at = ? WHERE identifier = ?",
		title,
		hash,
		now,
		identifier,
	)
	errors.PanicOnError(e)
}
