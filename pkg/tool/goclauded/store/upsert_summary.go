package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"

func (s *Store) UpsertSummary(
	sessionIdentifier string,
	name string,
	body string,
) error {
	var existing summary.Summary
	result := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&existing)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return s.database.Model(&existing).Updates(
			map[string]any{
				"name": name,
				"body": body,
			},
		).Error
	}

	return s.database.Create(summary.New(sessionIdentifier, name, body)).Error
}
