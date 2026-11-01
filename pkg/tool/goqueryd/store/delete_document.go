package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Store) DeleteDocument(
	collection string,
	path string,
) (bool, error) {
	var hash string
	e := s.database.QueryRow(
		"SELECT hash FROM document WHERE collection = ? AND path = ?",
		collection,
		path,
	).Scan(&hash)

	if e == sql.ErrNoRows {
		return false, nil
	}

	if e != nil {
		return false, e
	}

	_, e = s.database.Exec(
		"DELETE FROM document WHERE collection = ? AND path = ?",
		collection,
		path,
	)

	if e != nil {
		return false, e
	}

	var otherRefs int
	errors.PanicOnError(
		s.database.QueryRow(
			"SELECT COUNT(*) FROM document WHERE hash = ?",
			hash,
		).Scan(&otherRefs),
	)

	if otherRefs == 0 {
		errors.PanicOnError(s.DeleteEmbeddings(hash))
		_, e = s.database.Exec("DELETE FROM content WHERE hash = ?", hash)
		errors.PanicOnError(e)
	}

	return true, nil
}
