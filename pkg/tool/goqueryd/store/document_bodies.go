package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
)

func (s *Store) DocumentBodies(collection string) ([]*record.DocumentBody, error) {
	sql := `
		SELECT d.collection, d.path, d.hash, c.body
		FROM document d
		JOIN content c ON d.hash = c.hash
		WHERE d.active = 1`
	var arguments []any

	if collection != "" {
		sql = join.Empty(sql, " AND d.collection = ?")
		arguments = append(arguments, collection)
	}

	rows, e := s.database.Query(sql, arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []*record.DocumentBody

	for rows.Next() {
		var d record.DocumentBody

		if f := rows.Scan(&d.Collection, &d.Path, &d.Hash, &d.Body); f != nil {
			return nil, f
		}

		result = append(result, &d)
	}

	return result, rows.Err()
}
