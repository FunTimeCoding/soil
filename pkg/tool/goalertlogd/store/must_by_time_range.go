package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Store) MustByTimeRange(
	start time.Time,
	end time.Time,
) []Record {
	result, e := s.ByTimeRange(start, end)
	errors.PanicOnError(e)

	return result
}
