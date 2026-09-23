package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/cruise"
	"gorm.io/gorm/clause"
	"time"
)

func (s *Store) SetCruise(
	session string,
	mode constant.Cruise,
	pace time.Duration,
) {
	errors.PanicOnError(
		s.mapper.Clauses(
			clause.OnConflict{
				Columns: []clause.Column{{Name: "session"}},
				DoUpdates: clause.Assignments(
					map[string]any{
						"mode":       mode,
						"pace":       pace,
						"updated_at": s.clock(),
					},
				),
			},
		).Create(cruise.New(session, mode, pace)).Error,
	)
}
