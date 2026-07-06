package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustByName(name string) []Record {
	result, e := s.ByName(name)
	errors.PanicOnError(e)

	return result
}
