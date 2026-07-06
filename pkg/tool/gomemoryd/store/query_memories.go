package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Store) queryMemories(
	memoryType string,
	tag string,
	activeOnly bool,
	rootsOnly *bool,
) ([]MemorySummary, error) {
	var parts []string
	var arguments []any
	parts = append(
		parts,
		`SELECT DISTINCT m.identifier, m.name, m.description, m.type, m.updated_at, m.parent_identifier
		FROM memory m`,
	)

	if tag != "" {
		parts = append(
			parts,
			`JOIN memory_tag t ON t.memory_identifier = m.identifier AND t.tag = ?`,
		)
		arguments = append(arguments, tag)
	}

	parts = append(parts, `WHERE 1=1`)

	if activeOnly {
		parts = append(parts, `AND m.is_active = 1`)
	}

	if memoryType != "" {
		parts = append(parts, `AND m.type = ?`)
		arguments = append(arguments, memoryType)
	}

	if rootsOnly != nil && *rootsOnly {
		parts = append(parts, `AND m.parent_identifier IS NULL`)
	}

	parts = append(parts, `ORDER BY m.updated_at DESC`)
	rows, e := s.database.Query(join.Space(parts...), arguments...)

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
