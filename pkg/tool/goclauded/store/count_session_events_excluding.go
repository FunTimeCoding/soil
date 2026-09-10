package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"

func (s *Store) CountSessionEventsExcluding(
	sessionIdentifier string,
	kinds []string,
) (int64, error) {
	var result int64

	return result, s.database.Model(event.Stub()).Where(
		"session_identifier = ? AND kind NOT IN ?",
		sessionIdentifier,
		kinds,
	).Count(&result).Error
}
