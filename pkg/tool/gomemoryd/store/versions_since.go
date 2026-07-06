package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) VersionsSince(
	since string,
	limit int,
	offset int,
) ([]Version, error) {
	rows, e := s.database.Query(
		`SELECT identifier, memory_identifier, name, content, description, changed_at, change_type, source
		FROM memory_version WHERE changed_at > ? ORDER BY identifier LIMIT ? OFFSET ?`,
		since,
		limit,
		offset,
	)

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
