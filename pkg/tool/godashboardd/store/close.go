package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Close() {
	d, e := s.mapper.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(d.Close())
}
