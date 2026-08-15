package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"

func (s *Store) MarkCount() (int64, error) {
	var result int64
	e := s.database.Model(mark.Stub()).Count(&result).Error

	return result, e
}
