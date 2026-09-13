package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustDelete(
	callsign string,
	root string,
) int {
	result, e := s.Delete(callsign, root)
	errors.PanicOnError(e)

	return result
}
