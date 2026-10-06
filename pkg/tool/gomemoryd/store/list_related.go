package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Store) ListRelated(identifier int64) ([]record.Related, error) {
	rows, e := s.database.Query(
		`SELECT m.identifier, m.name, m.description, m.scope, r.type
		FROM memory_relation r
		JOIN memory m ON m.identifier = CASE
			WHEN r.source_identifier = ? THEN r.target_identifier
			ELSE r.source_identifier
		END
		WHERE (r.source_identifier = ? OR r.target_identifier = ?)
		AND m.is_active = 1
		ORDER BY r.created_at`,
		identifier,
		identifier,
		identifier,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []record.Related

	for rows.Next() {
		var r record.Related
		f := rows.Scan(
			&r.Identifier,
			&r.Name,
			&r.Description,
			&r.Scope,
			&r.Type,
		)

		if f != nil {
			return nil, f
		}

		result = append(result, r)
	}

	for i := range result {
		result[i].Tags = s.tagsForMemory(result[i].Identifier)
	}

	return result, nil
}
