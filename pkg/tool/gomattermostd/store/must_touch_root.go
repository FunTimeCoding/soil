package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Store) MustTouchRoot(root string, at time.Time) {
	errors.PanicOnError(s.TouchRoot(root, at))
}
