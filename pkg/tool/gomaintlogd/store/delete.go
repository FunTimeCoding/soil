package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) Delete(identifier uint) error {
	result := s.database.Delete(entry.New(), identifier)

	if result.Error != nil {
		return fmt.Errorf("delete entry: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return not_found.New("entry", identifier)
	}

	return nil
}
