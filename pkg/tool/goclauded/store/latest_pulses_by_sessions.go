package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) LatestPulsesBySessions(
	sessionIdentifiers []string,
) (map[string]*pulse.Pulse, error) {
	result := map[string]*pulse.Pulse{}

	if len(sessionIdentifiers) == 0 {
		return result, nil
	}

	var rows []pulse.Pulse
	e := s.database.Where("session_identifier IN ?", sessionIdentifiers).Order(
		"created_at DESC, identifier DESC",
	).Find(
		&rows,
	).Error

	if e != nil {
		return nil, e
	}

	for i := range rows {
		if result[rows[i].SessionIdentifier] == nil {
			result[rows[i].SessionIdentifier] = &rows[i]
		}
	}

	return result, nil
}
