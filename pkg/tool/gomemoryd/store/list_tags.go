package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Store) ListTags() ([]record.TagCount, error) {
	rows, e := s.database.Query(
		`SELECT t.tag, COUNT(*) as count
		FROM memory_tag t
		JOIN memory m ON m.identifier = t.memory_identifier AND m.is_active = 1
		GROUP BY t.tag
		ORDER BY count DESC, t.tag`,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []record.TagCount

	for rows.Next() {
		var tc record.TagCount
		e := rows.Scan(&tc.Tag, &tc.Count)

		if e != nil {
			return nil, e
		}

		result = append(result, tc)
	}

	return result, nil
}
