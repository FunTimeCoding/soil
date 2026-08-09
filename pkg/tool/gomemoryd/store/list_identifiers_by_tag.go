package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) ListIdentifiersByTag(tag string) (map[int64]bool, error) {
	rows, e := s.database.Query(
		`SELECT memory_identifier FROM memory_tag WHERE tag = ?`,
		tag,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	result := map[int64]bool{}

	for rows.Next() {
		var identifier int64

		if f := rows.Scan(&identifier); f != nil {
			return nil, f
		}

		result[identifier] = true
	}

	return result, nil
}
