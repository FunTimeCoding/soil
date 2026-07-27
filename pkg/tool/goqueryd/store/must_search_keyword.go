package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustSearchKeyword(
	query string,
	limit int,
	collection string,
	full bool,
	metadata map[string]string,
) []SearchResult {
	result, e := s.SearchKeyword(query, limit, collection, full, metadata)
	errors.PanicOnError(e)

	return result
}
