package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/tracker_state"
)

func (s *Store) OrphanTrackerStates() ([]string, error) {
	var result []string

	return result, s.database.Model(tracker_state.Stub()).Where(
		"identifier NOT IN (?)",
		s.database.Model(session.Stub()).Select(constant.Identifier),
	).Order(constant.Identifier).Pluck(constant.Identifier, &result).Error
}
