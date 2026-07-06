package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustListDocuments(collection string) []SearchResult {
	results, e := s.ListDocuments(collection, nil, 0, 0, false)
	errors.PanicOnError(e)

	return results
}
