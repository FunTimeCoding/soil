package store

func (s *Store) GetMemory(identifier int64) (*Memory, error) {
	row := s.database.QueryRow(
		`SELECT identifier, name, content, description, type, scope, created_at, updated_at, is_active, parent_identifier,
			provenance_file, provenance_anchor, provenance_hash, ordinal
		FROM memory WHERE identifier = ?`,
		identifier,
	)
	m := &Memory{}
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
