package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) updateDocumentTitle(
	identifier int,
	title string,
	now string,
) {
	_, e := s.database.Exec(
		"UPDATE document SET title = ?, modified_at = ? WHERE identifier = ?",
		title,
		now,
		identifier,
	)
	errors.PanicOnError(e)
}
