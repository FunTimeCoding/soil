package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/impression_outbox"

func (s *Store) DeleteImpression(identifier uint) error {
	return s.database.Delete(impression_outbox.Stub(), identifier).Error
}
