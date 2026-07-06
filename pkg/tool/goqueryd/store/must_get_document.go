package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func (s *Store) MustGetDocument(reference string) *result.Document {
	result, e := s.GetDocument(reference)
	errors.PanicOnError(e)

	return result
}
