package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) InsertEmbedding(
	hash string,
	sequence int,
	position int,
	vector []float32,
	now string,
) {
	blob := float32ToBytes(vector)
	_, e := s.database.Exec(
		`INSERT OR REPLACE INTO embedding
		(hash, sequence, position, vector, model, embedded_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		hash,
		sequence,
		position,
		blob,
		s.model,
		now,
	)
	errors.PanicOnError(e)
}
