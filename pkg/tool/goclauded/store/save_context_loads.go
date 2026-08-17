package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"gorm.io/gorm/clause"
)

func (s *Store) SaveContextLoads(loads []context_load.Load) error {
	if len(loads) == 0 {
		return nil
	}

	return s.database.Clauses(clause.OnConflict{DoNothing: true}).Create(
		&loads,
	).Error
}
