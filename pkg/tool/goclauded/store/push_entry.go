package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) PushEntry(e *queue.Entry) error {
	return s.database.Create(e).Error
}
