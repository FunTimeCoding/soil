package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func (s *Store) MustBySerial(serial string) *record.Record {
	result, e := s.BySerial(serial)
	errors.PanicOnError(e)

	return result
}
