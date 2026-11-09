package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
)

func (s *Store) MustByTimeRange(w *window.Window) []record.Record {
	result, e := s.ByTimeRange(w)
	errors.PanicOnError(e)

	return result
}
