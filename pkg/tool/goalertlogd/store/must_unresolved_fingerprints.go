package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustUnresolvedFingerprints() []string {
	result, e := s.UnresolvedFingerprints()
	errors.PanicOnError(e)

	return result
}
