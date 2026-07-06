package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustUnresolved() []UnresolvedRecord {
	result, e := s.Unresolved()
	errors.PanicOnError(e)

	return result
}
