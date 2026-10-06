package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
)

func (s *Store) MustListDocuments(collection string) []search.Result {
	results, e := s.ListDocuments(collection, nil, 0, 0, false)
	errors.PanicOnError(e)

	return results
}
