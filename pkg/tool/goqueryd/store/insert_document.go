package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) InsertDocument(
	collection string,
	path string,
	title string,
	hash string,
	now string,
) {
	_, e := s.database.Exec(
		`INSERT INTO document (collection, path, title, hash, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(collection, path) DO UPDATE SET
			title = excluded.title,
			hash = excluded.hash,
			modified_at = excluded.modified_at`,
		collection,
		path,
		title,
		hash,
		now,
		now,
	)
	errors.PanicOnError(e)
}
