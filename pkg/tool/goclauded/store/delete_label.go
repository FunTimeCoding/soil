package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Store) DeleteLabel(
	sessionIdentifier string,
	key string,
) (string, error) {
	var existing label.Label
	result := s.database.Where(
		"session_identifier = ? AND key = ?",
		sessionIdentifier,
		key,
	).Limit(1).Find(&existing)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", nil
	}

	return existing.Value, s.database.Delete(&existing).Error
}
