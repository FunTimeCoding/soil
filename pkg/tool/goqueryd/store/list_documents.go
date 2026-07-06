package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
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

	i := 0

	for key, value := range metadata {
		alias := fmt.Sprintf("m%d", i)
		parts = append(
			parts,
			fmt.Sprintf(
				`JOIN document_metadata %s
				ON %s.document_identifier = d.identifier
				AND %s.key = ? AND %s.value = ?`,
				alias,
				alias,
				alias,
				alias,
			),
		)
		arguments = append(arguments, key, value)
		i++
	}

	parts = append(parts, `WHERE d.collection = ? AND d.active = 1`)
	arguments = append(arguments, collection)
	parts = append(parts, `ORDER BY d.modified_at DESC`)

	if limit > 0 {
		parts = append(
			parts,
			fmt.Sprintf(`LIMIT %d OFFSET %d`, limit, offset),
		)
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
		r.FilePath = join.Empty(r.Collection, separator.Slash, r.Path)
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
