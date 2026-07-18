package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustSaveChannel(
	identifier int64,
	name string,
) {
	errors.PanicOnError(s.SaveChannel(identifier, name))
}
