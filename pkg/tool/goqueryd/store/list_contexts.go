package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func (s *Store) ListContexts() []result.ContextEntry {
	rows, e := s.database.Query(
		"SELECT collection, path_prefix, description FROM context ORDER BY collection, path_prefix",
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var r []result.ContextEntry

	for rows.Next() {
		var c result.ContextEntry
		errors.PanicOnError(
			rows.Scan(
				&c.Collection,
				&c.PathPrefix,
				&c.Description,
			),
		)
		r = append(r, c)
	}

	errors.PanicOnError(rows.Err())

	return r
}
