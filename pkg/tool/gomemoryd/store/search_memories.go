package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Store) SearchMemories(
	query string,
	limit int,
	memoryType string,
	tag string,
) ([]SearchResult, error) {
	var parts []string
	var arguments []any
	parts = append(
		parts,
		`SELECT m.identifier, m.name, m.content, m.description, m.type, m.updated_at, rank
		FROM memory_full_text_search f
		JOIN memory m ON m.identifier = f.rowid`,
	)

	if tag != "" {
		parts = append(
			parts,
			`JOIN memory_tag t ON t.memory_identifier = m.identifier AND t.tag = ?`,
		)
		arguments = append(arguments, tag)
	}

	parts = append(
		parts,
		`WHERE memory_full_text_search MATCH ? AND m.is_active = 1`,
	)
	arguments = append(arguments, query)

	if memoryType != "" {
		parts = append(parts, `AND m.type = ?`)
		arguments = append(arguments, memoryType)
	}

	parts = append(parts, fmt.Sprintf(`ORDER BY rank LIMIT %d`, limit))
	rows, e := s.database.Query(join.Space(parts...), arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []SearchResult

	for rows.Next() {
		var r SearchResult
		e := rows.Scan(
			&r.Identifier,
			&r.Name,
			&r.Content,
			&r.Description,
			&r.Type,
			&r.UpdatedAt,
			&r.Rank,
		)

		if e != nil {
			return nil, e
		}

		r.Tags = s.tagsForMemory(r.Identifier)
		result = append(result, r)
	}

	return result, nil
}
