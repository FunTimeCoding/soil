package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) PendingPulses(sessionIdentifier string) ([]pulse.Pulse, error) {
	var result []pulse.Pulse

	if e := s.database.Where(
		"session_identifier = ? AND from_name = '' AND consumed = ?",
		sessionIdentifier,
		false,
	).Order("created_at").Find(&result).Error; e != nil {
		return nil, e
	}

	if len(result) > 0 {
		if e := s.database.Model(pulse.Stub()).Where(
			"session_identifier = ? AND from_name = '' AND consumed = ?",
			sessionIdentifier,
			false,
		).Update("consumed", true).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
