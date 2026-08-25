package store

func (s *Store) DeleteRelation(
	sourceIdentifier int64,
	targetIdentifier int64,
) (bool, error) {
	result, e := s.database.Exec(
		`DELETE FROM memory_relation
		WHERE source_identifier = ? AND target_identifier = ?`,
		sourceIdentifier,
		targetIdentifier,
	)

	if e != nil {
		return false, e
	}

	affected, e := result.RowsAffected()

	if e != nil {
		return false, e
	}

	return affected > 0, nil
}
