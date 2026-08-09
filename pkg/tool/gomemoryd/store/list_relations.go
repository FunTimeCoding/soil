package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) ListRelations() ([]RelationOverview, error) {
	rows, e := s.database.Query(
		`SELECT r.source_identifier, s.name, s.scope,
			r.target_identifier, t.name, t.scope,
			r.type, r.created_at
		FROM memory_relation r
		JOIN memory s ON s.identifier = r.source_identifier
		JOIN memory t ON t.identifier = r.target_identifier
		WHERE s.is_active = 1 AND t.is_active = 1
		ORDER BY s.scope, s.name, t.scope, t.name`,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []RelationOverview

	for rows.Next() {
		var r RelationOverview
		f := rows.Scan(
			&r.SourceIdentifier,
			&r.SourceName,
			&r.SourceScope,
			&r.TargetIdentifier,
			&r.TargetName,
			&r.TargetScope,
			&r.Type,
			&r.CreatedAt,
		)

		if f != nil {
			return nil, f
		}

		result = append(result, r)
	}

	return result, nil
}
