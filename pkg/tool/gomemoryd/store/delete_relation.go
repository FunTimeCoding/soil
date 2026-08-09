package store

func (s *Store) DeleteRelation(
	sourceID int64,
	targetID int64,
) (bool, error) {
	result, e := s.database.Exec(
		`DELETE FROM memory_relation
		WHERE source_identifier = ? AND target_identifier = ?`,
		sourceID,
		targetID,
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
