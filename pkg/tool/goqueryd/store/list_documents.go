package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Store) ListDocuments(
	collection string,
	metadata map[string]string,
	limit int,
	offset int,
	full bool,
) ([]SearchResult, error) {
	var parts []string
	var arguments []any

	if full {
		parts = append(
			parts,
			`SELECT d.identifier, d.collection, d.path, d.title, d.hash, d.modified_at, c.body
			FROM document d
			JOIN content c ON c.hash = d.hash`,
		)
	} else {
		parts = append(
			parts,
			`SELECT d.identifier, d.collection, d.path, d.title, d.hash, d.modified_at, ''
			FROM document d`,
		)
	}

	joins, joinArguments := metadataJoins(metadata)

	if joins != "" {
		parts = append(parts, joins)
		arguments = append(arguments, joinArguments...)
	}

	parts = append(parts, `WHERE d.collection = ? AND d.active = 1`)
	arguments = append(arguments, collection)
	parts = append(parts, `ORDER BY d.modified_at DESC`)

	if limit > 0 {
		parts = append(parts, fmt.Sprintf(`LIMIT %d OFFSET %d`, limit, offset))
	}

	rows, e := s.database.Query(join.Space(parts...), arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []SearchResult
	var identifiers []int

	for rows.Next() {
		var r SearchResult
		var identifier int
		var modifiedAt string

		if f := rows.Scan(
			&identifier,
			&r.Collection,
			&r.Path,
			&r.Title,
			&r.Hash,
			&modifiedAt,
			&r.Body,
		); f != nil {
			return nil, f
		}

		r.VirtualPath = buildVirtualPath(r.Collection, r.Path)
		r.FilePath = join.Empty(r.Collection, constant.Slash, r.Path)
		r.Source = "list"
		identifiers = append(identifiers, identifier)
		result = append(result, r)
	}

	if e := rows.Err(); e != nil {
		return nil, e
	}

	s.enrichMetadata(result, identifiers)

	return result, nil
}
