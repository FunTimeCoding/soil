package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) After(
	identifier uint,
	limit int,
) ([]record.Record, error) {
	var result []record.Record

	return result, s.database.
		Where("identifier > ?", identifier).
		Order("identifier").
		Limit(limit).
		Find(&result).Error
}
