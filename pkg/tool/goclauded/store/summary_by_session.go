package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"

func (s *Store) SummaryBySession(sessionIdentifier string) (string, error) {
	var i summary.Summary
	result := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&i)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", nil
	}

	return i.Body, nil
}
