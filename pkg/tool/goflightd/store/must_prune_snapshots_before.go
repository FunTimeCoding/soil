package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Store) MustPruneSnapshotsBefore(cutoff time.Time) int64 {
	count, e := s.PruneSnapshotsBefore(cutoff)
	errors.PanicOnError(e)

	return count
}
