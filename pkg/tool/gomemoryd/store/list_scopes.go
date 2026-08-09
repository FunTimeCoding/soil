package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) ListScopes() ([]string, error) {
	rows, e := s.database.Query(
		`SELECT DISTINCT scope FROM memory WHERE is_active = 1 ORDER BY scope`,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []string

	for rows.Next() {
		var scope string

		if f := rows.Scan(&scope); f != nil {
			return nil, f
		}

		result = append(result, scope)
	}

	return result, nil
}
