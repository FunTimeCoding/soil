package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustCount() int {
	result, e := s.Count()
	errors.PanicOnError(e)

	return result
}
