package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustResolve(fingerprint string) {
	errors.PanicOnError(s.Resolve(fingerprint))
}
