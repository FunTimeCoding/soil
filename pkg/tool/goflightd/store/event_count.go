package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"

func (s *Store) EventCount() (int64, error) {
	var result int64
	e := s.database.Model(event.Stub()).Count(&result).Error

	return result, e
}
