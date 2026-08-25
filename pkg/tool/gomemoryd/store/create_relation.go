package store

import "time"

func (s *Store) CreateRelation(
	sourceIdentifier int64,
	targetIdentifier int64,
	relationType string,
) error {
	now := time.Now().UTC().Format(time.RFC3339)
	_, e := s.database.Exec(
		`INSERT INTO memory_relation (source_identifier, target_identifier, created_at, type)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(source_identifier, target_identifier)
		DO UPDATE SET type = excluded.type WHERE excluded.type != ''`,
		sourceIdentifier,
		targetIdentifier,
		now,
		relationType,
	)

	return e
}
