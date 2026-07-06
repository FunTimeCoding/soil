package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Store) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
) error {
	return s.database.Create(
		pulse.New(sessionIdentifier, fromName, body),
	).Error
}
