package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"gorm.io/gorm/clause"
	"time"
)

func (s *Store) SetCruise(
	sessionIdentifier string,
	mode constant.Cruise,
	pace time.Duration,
) error {
	v := cruise_setting.New(sessionIdentifier, mode, pace)
	v.UpdatedAt = s.clock().UTC()

	return s.database.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "session_identifier"}},
			DoUpdates: clause.Assignments(
				map[string]any{
					"mode":       mode,
					"pace":       pace,
					"updated_at": v.UpdatedAt,
				},
			),
		},
	).Create(v).Error
}
