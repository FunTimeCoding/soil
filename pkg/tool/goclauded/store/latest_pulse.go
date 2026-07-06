package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) LatestPulse(sessionIdentifier string) (*pulse.Pulse, error) {
	var result pulse.Pulse
	r := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Order("created_at DESC").Limit(1).Find(&result)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}

	return &result, nil
}
