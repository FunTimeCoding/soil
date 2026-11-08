package store

func (s *Store) AddLabel(
	sessionIdentifier string,
	key string,
	value string,
) ([]string, error) {
	old, e := s.GetLabel(sessionIdentifier, key)

	if e != nil {
		return nil, e
	}

	return old, s.database.Exec(
		`INSERT INTO label (session_identifier, key, value)
		SELECT ?, ?, ? WHERE NOT EXISTS (
			SELECT 1 FROM label
			WHERE session_identifier = ? AND key = ? AND value = ?
		)`,
		sessionIdentifier,
		key,
		value,
		sessionIdentifier,
		key,
		value,
	).Error
}
