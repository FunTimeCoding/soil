package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) ListDocumentSourced(scope string) ([]SourcedMemory, error) {
	rows, e := s.database.Query(
		`SELECT identifier, name, parent_identifier, provenance_file, provenance_anchor, provenance_hash, ordinal
		FROM memory
		WHERE scope = ? AND provenance_file != '' AND is_active = 1
		ORDER BY provenance_file, ordinal`,
		scope,
	)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(rows)
	var result []SourcedMemory

	for rows.Next() {
		var m SourcedMemory
		e := rows.Scan(
			&m.Identifier,
			&m.Name,
			&m.ParentIdentifier,
			&m.ProvenanceFile,
			&m.ProvenanceAnchor,
			&m.ProvenanceHash,
			&m.Ordinal,
		)

		if e != nil {
			return nil, e
		}

		result = append(result, m)
	}

	return result, nil
}
