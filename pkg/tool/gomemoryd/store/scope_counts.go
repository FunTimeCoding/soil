package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Store) ScopeCounts() ([]record.ScopeCount, error) {
	rows, e := s.database.Query(
		`SELECT scope, COUNT(*) FROM memory WHERE is_active = 1
		GROUP BY scope ORDER BY COUNT(*) DESC, scope`,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []record.ScopeCount

	for rows.Next() {
		var one record.ScopeCount

		if f := rows.Scan(&one.Scope, &one.Count); f != nil {
			return nil, f
		}

		result = append(result, one)
	}

	return result, rows.Err()
}
