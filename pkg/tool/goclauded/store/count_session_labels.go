package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Store) CountSessionLabels(sessionIdentifier string) (int64, error) {
	var result int64

	return result, s.database.Model(label.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&result).Error
}
