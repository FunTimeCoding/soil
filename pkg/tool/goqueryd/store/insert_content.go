package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) InsertContent(
	hash string,
	content string,
	now string,
) {
	_, e := s.database.Exec(
		"INSERT OR IGNORE INTO content (hash, body, created_at) VALUES (?, ?, ?)",
		hash,
		content,
		now,
	)
	errors.PanicOnError(e)
}
