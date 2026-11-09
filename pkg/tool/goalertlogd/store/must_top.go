package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
)

func (s *Store) MustTop(
	n int,
	w *window.Window,
) []top.Record {
	result, e := s.Top(n, w)
	errors.PanicOnError(e)

	return result
}
