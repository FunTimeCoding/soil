package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) Close() {
	errors.PanicOnError(s.database.Close())
}
