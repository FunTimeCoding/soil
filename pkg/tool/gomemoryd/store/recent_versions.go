package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Store) RecentVersions(
	since string,
	limit int,
	excludeTag string,
) ([]Version, error) {
	parts := []string{
		`SELECT identifier, memory_identifier, name, content, description, changed_at, change_type, source
		FROM memory_version WHERE changed_at > ?`,
	}
	arguments := []any{since}

	if excludeTag != "" {
		parts = append(
			parts,
			`AND memory_identifier NOT IN (SELECT memory_identifier FROM memory_tag WHERE tag = ?)`,
		)
		arguments = append(arguments, excludeTag)
	}

	parts = append(parts, `ORDER BY identifier DESC LIMIT ?`)
	arguments = append(arguments, limit)
	rows, e := s.database.Query(join.Space(parts...), arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var versions []Version

	for rows.Next() {
		var v Version
		e := rows.Scan(
			&v.Identifier,
			&v.MemoryIdentifier,
			&v.Name,
			&v.Content,
			&v.Description,
			&v.ChangedAt,
			&v.ChangeType,
			&v.Source,
		)

		if e != nil {
			return nil, e
		}

		versions = append(versions, v)
	}

	return versions, nil
}
