package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) PushQueueImmediate(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) error {
	return s.database.Create(
		queue.NewImmediate(sessionIdentifier, callsign, kind, body),
	).Error
}
