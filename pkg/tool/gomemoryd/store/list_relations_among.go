package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Store) ListRelationsAmong(
	identifiers []int64,
) ([]RelationOverview, error) {
	if len(identifiers) == 0 {
		return nil, nil
	}

	placeholders := make([]string, 0, len(identifiers))
	var arguments []any

	for _, identifier := range identifiers {
		placeholders = append(placeholders, constant.Question)
		arguments = append(arguments, identifier)
	}

	set := join.CommaSpace(placeholders)
	arguments = append(arguments, arguments...)
	rows, e := s.database.Query(
		fmt.Sprintf(
			`SELECT r.source_identifier, s.name, s.scope,
				r.target_identifier, t.name, t.scope,
				r.type, r.created_at
			FROM memory_relation r
			JOIN memory s ON s.identifier = r.source_identifier
			JOIN memory t ON t.identifier = r.target_identifier
			WHERE (r.source_identifier IN (%s)
				OR r.target_identifier IN (%s))
			AND s.is_active = 1 AND t.is_active = 1
			ORDER BY s.name, t.name`,
			set,
			set,
		),
		arguments...,
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
