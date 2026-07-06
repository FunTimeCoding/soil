package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) listMemoriesWithParent(parentID int64) ([]MemorySummary, error) {
	rows, e := s.database.Query(
		`SELECT identifier, name, description, type, updated_at, parent_identifier
		FROM memory
		WHERE parent_identifier = ? AND is_active = 1
		ORDER BY updated_at DESC`,
		parentID,
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
