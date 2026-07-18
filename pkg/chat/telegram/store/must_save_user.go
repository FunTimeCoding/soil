package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustSaveUser(
	identifier int64,
	name string,
) {
	errors.PanicOnError(s.SaveUser(identifier, name))
}
