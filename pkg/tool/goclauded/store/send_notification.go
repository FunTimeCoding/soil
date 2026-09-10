package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"

func (s *Store) SendNotification(
	sessionIdentifier string,
	callsign string,
	source string,
	body string,
) error {
	return s.database.Create(
		notification.New(sessionIdentifier, callsign, source, body),
	).Error
}
