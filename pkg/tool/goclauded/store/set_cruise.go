package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"gorm.io/gorm/clause"
	"time"
)

func (s *Store) SetCruise(
	sessionIdentifier string,
	grace *time.Duration,
) error {
	v := cruise_setting.New(sessionIdentifier, grace)
	v.UpdatedAt = s.clock()

	return s.database.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "session_identifier"}},
			DoUpdates: clause.Assignments(
				map[string]any{
					constant.GraceColumn: grace,
					"updated_at":         v.UpdatedAt,
				},
			),
		},
	).Create(v).Error
}
