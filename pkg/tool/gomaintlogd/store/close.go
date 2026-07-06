package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) Close() {
	d, e := s.database.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(d.Close())
}
