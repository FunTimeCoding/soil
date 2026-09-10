package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"

func (s *Store) CountSessionEvents(sessionIdentifier string) (int64, error) {
	var result int64

	return result, s.database.Model(event.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&result).Error
}
