package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) InsertEmbeddingFailure(
	hash string,
	sequence int,
	position int,
	reason string,
	now string,
) {
	_, e := s.database.Exec(
		`INSERT OR REPLACE INTO embedding_failure
		(hash, sequence, position, reason, failed_at)
		VALUES (?, ?, ?, ?, ?)`,
		hash,
		sequence,
		position,
		reason,
		now,
	)
	errors.PanicOnError(e)
}
