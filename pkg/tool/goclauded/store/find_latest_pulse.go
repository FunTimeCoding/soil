package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) FindLatestPulse(
	sessionIdentifier string,
) (*pulse.Pulse, bool, error) {
	var result pulse.Pulse
	r := s.database.Where("session_identifier = ?", sessionIdentifier).Order(
		"created_at DESC",
	).Limit(
		1,
	).Find(
		&result,
	)

	if r.Error != nil {
		return nil, false, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, false, nil
	}

	return &result, true, nil
}
