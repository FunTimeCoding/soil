package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"time"
)

func (s *Store) MustByTimeRange(
	start time.Time,
	end time.Time,
) []record.Record {
	result, e := s.ByTimeRange(start, end)
	errors.PanicOnError(e)

	return result
}
