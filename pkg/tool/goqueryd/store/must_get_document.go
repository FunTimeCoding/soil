package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func (s *Store) MustGetDocument(reference string) *result.Document {
	result, found, e := s.FindDocument(reference)
	errors.PanicOnError(e)

	if !found {
		panic(not_found.New("document", reference))
	}

	return result
}
