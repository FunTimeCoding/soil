package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
)

func (s *Store) MustCreate(r record.Record) {
	errors.PanicOnError(s.Create(r))
}
