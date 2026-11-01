package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) Prune() int64 {
	_, e := s.database.Exec(
		`DELETE FROM embedding
		WHERE NOT EXISTS (SELECT 1 FROM document d WHERE d.hash = embedding.hash)`,
	)
	errors.PanicOnError(e)
	_, e = s.database.Exec(
		`DELETE FROM embedding_failure
		WHERE NOT EXISTS (
			SELECT 1 FROM document d WHERE d.hash = embedding_failure.hash
		)`,
	)
	errors.PanicOnError(e)
	r, e := s.database.Exec(
		`DELETE FROM content
		WHERE NOT EXISTS (SELECT 1 FROM document d WHERE d.hash = content.hash)`,
	)
	errors.PanicOnError(e)
	result, f := r.RowsAffected()
	errors.PanicOnError(f)

	return result
}
