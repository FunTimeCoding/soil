package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) ListRelated(identifier int64) ([]Related, error) {
	rows, e := s.database.Query(
		`SELECT m.identifier, m.name, m.description
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
	var result []Related

	for rows.Next() {
		var r Related
		f := rows.Scan(&r.Identifier, &r.Name, &r.Description)

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
