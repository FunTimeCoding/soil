package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) Update(v *entry.Entry) error {
	if e := s.database.Save(v).Error; e != nil {
		return fmt.Errorf("update entry: %w", e)
	}

	return nil
}
