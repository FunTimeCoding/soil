package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/tracker_state"

func (s *Store) DeleteTrackerState(identifier string) error {
	return s.database.Where(
		"identifier = ?",
		identifier,
	).Delete(tracker_state.Stub()).Error
}
