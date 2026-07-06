package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) PushQueue(
	callsign string,
	kind string,
	body string,
) error {
	return s.database.Create(queue.New(callsign, kind, body)).Error
}
