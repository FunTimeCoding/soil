package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustResolve(key string) {
	errors.PanicOnError(s.Resolve(key))
}
