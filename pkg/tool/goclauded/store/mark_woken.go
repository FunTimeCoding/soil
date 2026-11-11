package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Store) MarkWoken(sessionIdentifier string) error {
	now := s.clock()
	v := cruise_setting.New(sessionIdentifier, nil)
	v.Woken = 1
	v.WokenAt = &now
	v.UpdatedAt = now

	return s.database.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "session_identifier"}},
			DoUpdates: clause.Assignments(
				map[string]any{
					"woken":    gorm.Expr("woken + 1"),
					"woken_at": now,
				},
			),
		},
	).Create(v).Error
}
