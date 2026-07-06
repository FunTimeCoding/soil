package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

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
		constant.EmbedModel,
		now,
	)
	errors.PanicOnError(e)
}
