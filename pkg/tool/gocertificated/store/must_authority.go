package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func (s *Store) MustAuthority(name string) *record.Record {
	result, e := s.Authority(name)
	errors.PanicOnError(e)

	return result
}
