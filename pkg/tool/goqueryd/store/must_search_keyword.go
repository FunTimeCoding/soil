package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
)

func (s *Store) MustSearchKeyword(
	query string,
	limit int,
	collection string,
	full bool,
	metadata map[string]string,
) []search.Result {
	result, e := s.SearchKeyword(query, limit, collection, full, metadata)
	errors.PanicOnError(e)

	return result
}
