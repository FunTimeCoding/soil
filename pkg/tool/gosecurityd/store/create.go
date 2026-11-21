package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"gorm.io/gorm/clause"
)

func (s *Store) Create(v *flag.Flag) error {
	r := s.database.Clauses(clause.OnConflict{DoNothing: true}).Create(v)

	if r.Error != nil {
		return fmt.Errorf("create flag: %w", r.Error)
	}

	if r.RowsAffected == 0 {
		return conflict.Exists(constant.FlagKind, v.Subject)
	}

	return nil
}
