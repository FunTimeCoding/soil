package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) DeletePendingQueue(
	sessionIdentifier string,
	callsign string,
	kind string,
) error {
	condition, arguments := sessionKeyMatch(sessionIdentifier, callsign)

	return s.database.Where(
		condition,
		arguments...,
	).Where(
		"kind = ? AND consumed_at IS NULL",
		kind,
	).Delete(queue.Stub()).Error
}
