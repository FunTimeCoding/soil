package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) PulsesBySession(sessionIdentifier string) ([]pulse.Pulse, error) {
	var result []pulse.Pulse

	return result, s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Order("created_at").Find(&result).Error
}
