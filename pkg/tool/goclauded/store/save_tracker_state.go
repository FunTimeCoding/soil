package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/tracker_state"
	"gorm.io/gorm/clause"
)

func (s *Store) SaveTrackerState(
	identifier string,
	state *tracker.State,
) {
	errors.PanicOnError(
		s.database.Clauses(clause.OnConflict{UpdateAll: true}).
			Create(
				tracker_state.New(
					identifier,
					string(notation.Marshal(state)),
					s.clock(),
				),
			).Error,
	)
}
