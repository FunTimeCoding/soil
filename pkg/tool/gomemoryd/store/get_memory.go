package store

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Store) GetMemory(identifier int64) (*record.Memory, error) {
	row := s.database.QueryRow(
		`SELECT identifier, name, content, description, type, scope, created_at, updated_at, is_active, parent_identifier,
			provenance_file, provenance_anchor, provenance_hash, ordinal
		FROM memory WHERE identifier = ?`,
		identifier,
	)
	m := record.NewMemory()
	var active int
	e := row.Scan(
		&m.Identifier,
		&m.Name,
		&m.Content,
		&m.Description,
		&m.Type,
		&m.Scope,
		&m.CreatedAt,
		&m.UpdatedAt,
		&active,
		&m.ParentIdentifier,
		&m.ProvenanceFile,
		&m.ProvenanceAnchor,
		&m.ProvenanceHash,
		&m.Ordinal,
	)

	if e != nil {
		return nil, e
	}

	m.IsActive = active == 1
	m.Tags = s.tagsForMemory(identifier)
	m.Metadata = s.metadataForMemory(identifier)

	return m, nil
}
