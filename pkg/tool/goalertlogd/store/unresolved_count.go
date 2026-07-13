package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) UnresolvedCount() (int, error) {
	var result int64
	e := s.database.
		Model(record.Stub()).
		Where("ended_at IS NULL").
		Count(&result).Error

	return int(result), e
}
