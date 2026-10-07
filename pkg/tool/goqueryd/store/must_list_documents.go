package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func (s *Store) MustListDocuments(collection string) []result.Search {
	results, e := s.ListDocuments(collection, nil, 0, 0, false)
	errors.PanicOnError(e)

	return results
}
