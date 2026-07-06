package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) DeletePendingQueue(
	callsign string,
	kind string,
) error {
	return s.database.Where(
		"callsign = ? AND kind = ? AND consumed = ?",
		callsign,
		kind,
		false,
	).Delete(queue.Stub()).Error
}
