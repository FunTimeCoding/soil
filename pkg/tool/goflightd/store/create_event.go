package store

import "github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"

func (s *Store) CreateEvent(v event.Event) error {
	return s.database.Create(&v).Error
}
