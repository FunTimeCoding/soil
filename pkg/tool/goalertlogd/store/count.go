package store

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Store) Count() (int, error) {
	var result int64
	e := s.database.Model(record.Stub()).Count(&result).Error

	return int(result), e
}
