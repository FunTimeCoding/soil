package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"time"
)

func (s *Store) PruneEventsBefore(cutoff time.Time) (int64, error) {
	result := s.database.Where("time < ?", cutoff).Delete(event.Stub())

	return result.RowsAffected, result.Error
}
