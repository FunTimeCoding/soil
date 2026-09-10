package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) CountSessionPulses(sessionIdentifier string) (int64, error) {
	var result int64

	return result, s.database.Model(pulse.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&result).Error
}
