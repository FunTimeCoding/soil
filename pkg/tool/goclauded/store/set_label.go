package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
)

func (s *Store) SetLabel(
	sessionIdentifier string,
	key string,
	value string,
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

	if result.RowsAffected > 0 {
		old := existing.Value

		if old == value {
			return old, nil
		}

		return old, s.database.Model(&existing).Update(
			constant.Value,
			value,
		).Error
	}

	return "", s.database.Create(
		label.New(sessionIdentifier, key, value),
	).Error
}
