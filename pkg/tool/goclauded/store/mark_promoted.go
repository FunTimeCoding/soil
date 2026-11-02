package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"gorm.io/gorm"
)

func (s *Store) MarkPromoted(sessionIdentifier string) error {
	return s.database.Model(cruise_setting.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Updates(
		map[string]any{
			"promoted_at": s.clock(),
			"woken":       gorm.Expr("woken + 1"),
		},
	).Error
}
