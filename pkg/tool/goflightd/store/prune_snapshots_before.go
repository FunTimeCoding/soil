package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"time"
)

func (s *Store) PruneSnapshotsBefore(cutoff time.Time) (int64, error) {
	result := s.database.Where("time < ?", cutoff).Delete(snapshot.Stub())

	return result.RowsAffected, result.Error
}
