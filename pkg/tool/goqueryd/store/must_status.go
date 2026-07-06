package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"
)

func (s *Store) MustStatus() *status.Status {
	r, e := s.Status()
	errors.PanicOnError(e)

	return r
}
