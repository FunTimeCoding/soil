package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
	"time"
)

func (s *Store) MustTop(
	n int,
	start time.Time,
	end time.Time,
) []top.Record {
	result, e := s.Top(n, start, end)
	errors.PanicOnError(e)

	return result
}
