package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) listMemoriesWithParent(parentIdentifier int64) ([]MemorySummary, error) {
	rows, e := s.database.Query(
		`SELECT identifier, name, description, type, scope, updated_at, parent_identifier
		FROM memory
		WHERE parent_identifier = ? AND is_active = 1
		ORDER BY ordinal, updated_at DESC`,
		parentIdentifier,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []MemorySummary

	for rows.Next() {
		var m MemorySummary
		e := rows.Scan(
			&m.Identifier,
			&m.Name,
			&m.Description,
			&m.Type,
			&m.Scope,
			&m.UpdatedAt,
			&m.ParentIdentifier,
		)

		if e != nil {
			return nil, e
		}

		m.Tags = s.tagsForMemory(m.Identifier)
		result = append(result, m)
	}

	return result, nil
}
