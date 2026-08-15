package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Store) MustPruneEventsBefore(cutoff time.Time) int64 {
	count, e := s.PruneEventsBefore(cutoff)
	errors.PanicOnError(e)

	return count
}
