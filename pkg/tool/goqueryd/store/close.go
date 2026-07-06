package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) Close() {
	errors.PanicClose(s.database)
}
