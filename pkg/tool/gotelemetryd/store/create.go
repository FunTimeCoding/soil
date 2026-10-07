package store

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"

func (s *Store) Create(e *usage_event.Event) error {
	return s.mapper.Create(e).Error
}
