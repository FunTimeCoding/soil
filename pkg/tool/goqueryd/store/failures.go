package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
)

func (s *Store) Failures(collection string) ([]*record.Failure, error) {
	sql := `
		SELECT d.collection, d.path, c.body, f.sequence, f.reason
		FROM embedding_failure f
		JOIN document d ON d.hash = f.hash AND d.active = 1
		JOIN content c ON c.hash = f.hash`
	var arguments []any

	if collection != "" {
		sql = join.Empty(sql, " WHERE d.collection = ?")
		arguments = append(arguments, collection)
	}

	sql = join.Empty(sql, " ORDER BY d.collection, d.path, f.sequence")
	rows, e := s.database.Query(sql, arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []*record.Failure

	for rows.Next() {
		var f record.Failure

		if g := rows.Scan(
			&f.Collection,
			&f.Path,
			&f.Body,
			&f.Sequence,
			&f.Reason,
		); g != nil {
			return nil, g
		}

		result = append(result, &f)
	}

	return result, rows.Err()
}
