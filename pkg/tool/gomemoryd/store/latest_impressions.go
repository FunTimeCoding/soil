package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Store) LatestImpressions(limit int) ([]record.Impression, error) {
	rows, e := s.database.Query(
		`SELECT identifier, content, source, created_at
		FROM impression ORDER BY identifier DESC LIMIT ?`,
		limit,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []record.Impression

	for rows.Next() {
		var i record.Impression
		e := rows.Scan(&i.Identifier, &i.Content, &i.Source, &i.CreatedAt)

		if e != nil {
			return nil, e
		}

		result = append(result, i)
	}

	return result, nil
}
