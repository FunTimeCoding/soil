package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Store) Labels(
	objectType string,
	objectIdentifier int32,
) ([]*label.Label, error) {
	var result []*label.Label

	return result, s.database.
		Where(
			"object_type = ? AND object_identifier = ?",
			objectType,
			objectIdentifier,
		).
		Order(constant.KeyColumn).
		Find(&result).Error
}
