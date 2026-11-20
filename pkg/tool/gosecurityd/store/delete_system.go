package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"gorm.io/gorm"
)

func (s *Store) DeleteSystem(v *system.System) error {
	return s.database.Transaction(
		func(t *gorm.DB) error {
			if e := t.Where(
				map[string]any{"system_identifier": v.Identifier},
			).Delete(installed.Stub()).Error; e != nil {
				return fmt.Errorf("delete packages: %w", e)
			}

			if f := t.Delete(v).Error; f != nil {
				return fmt.Errorf("delete system: %w", f)
			}

			return nil
		},
	)
}
