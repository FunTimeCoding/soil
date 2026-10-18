package store

import "time"

func (s *Store) CreateImpression(
	content string,
	source string,
	at time.Time,
) (int64, error) {
	result, e := s.database.Exec(
		`INSERT INTO impression (content, source, created_at) VALUES (?, ?, ?)`,
		content,
		source,
		at.Format(time.RFC3339),
	)

	if e != nil {
		return 0, e
	}

	return result.LastInsertId()
}
