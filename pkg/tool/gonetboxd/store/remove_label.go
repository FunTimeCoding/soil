package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Store) RemoveLabel(
	objectType string,
	objectIdentifier int32,
	key string,
) error {
	result := s.database.
		Where(
			"object_type = ? AND object_identifier = ? AND key = ?",
			objectType,
			objectIdentifier,
			key,
		).
		Delete(label.Stub())

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("label not found: %s", key)
	}

	return nil
}
