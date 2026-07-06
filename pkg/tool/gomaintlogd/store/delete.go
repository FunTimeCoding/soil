package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) Delete(identifier uint) error {
	if e := s.database.Delete(entry.New(), identifier).Error; e != nil {
		return fmt.Errorf("failed to delete entry: %w", e)
	}

	return nil
}
