package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) MustFindSimilarFiles(
	query string,
	limit int,
) []string {
	result, e := s.FindSimilarFiles(query, limit)
	errors.PanicOnError(e)

	return result
}
