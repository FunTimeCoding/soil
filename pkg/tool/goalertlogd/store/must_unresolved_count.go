package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustUnresolvedCount() int {
	result, e := s.UnresolvedCount()
	errors.PanicOnError(e)

	return result
}
