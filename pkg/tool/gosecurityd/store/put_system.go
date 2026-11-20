package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"gorm.io/gorm"
)

func (s *Store) PutSystem(
	v *system.System,
	packages []installed.Installed,
) error {
	return s.database.Transaction(
		func(t *gorm.DB) error {
			var existing system.System
			found, e := relational.FindOne(
				t.Where(map[string]any{"name": v.Name}),
				&existing,
			)

			if e != nil {
				return fmt.Errorf("find system: %w", e)
			}

			if found {
				v.Identifier = existing.Identifier
				v.CreatedAt = existing.CreatedAt
			}

			if f := t.Save(v).Error; f != nil {
				return fmt.Errorf("save system: %w", f)
			}

			if g := t.Where(
				map[string]any{"system_identifier": v.Identifier},
			).Delete(installed.Stub()).Error; g != nil {
				return fmt.Errorf("clear packages: %w", g)
			}

			for i := range packages {
				packages[i].Identifier = 0
				packages[i].SystemIdentifier = v.Identifier
			}

			if len(packages) == 0 {
				return nil
			}

			if h := t.CreateInBatches(packages, 500).Error; h != nil {
				return fmt.Errorf("save packages: %w", h)
			}

			return nil
		},
	)
}
