package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) ByName(name string) ([]record.Record, error) {
	var result []record.Record

	return result, s.database.
		Where("name = ?", name).
		Order("started_at DESC").
		Find(&result).Error
}
