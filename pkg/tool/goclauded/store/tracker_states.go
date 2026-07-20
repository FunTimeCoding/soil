package store

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/tracker_state"
)

func (s *Store) TrackerStates() map[string]*tracker.State {
	var records []tracker_state.Record
	errors.PanicOnError(s.database.Find(&records).Error)
	result := make(map[string]*tracker.State, len(records))

	for _, r := range records {
		state := tracker.New()

		if json.Unmarshal([]byte(r.Payload), state) != nil {
			continue
		}

		result[r.Identifier] = state
	}

	return result
}
