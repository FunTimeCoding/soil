package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Store) RecentImpressions(since string) ([]record.Impression, error) {
	rows, e := s.database.Query(
		`SELECT identifier, content, source, created_at
		FROM impression WHERE created_at > ? ORDER BY identifier DESC`,
		since,
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
