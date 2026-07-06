package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Store) documentIdentifiers(keys []documentKey) map[documentKey]int {
	if len(keys) == 0 {
		return nil
	}

	var arguments []any

	for _, k := range keys {
		arguments = append(arguments, k.Collection, k.Path)
	}

	rows, e := s.database.Query(
		fmt.Sprintf(
			`SELECT collection, path, identifier
			FROM document
			WHERE active = 1
			AND (%s)`,
			orPairs(len(keys)),
		),
		arguments...,
	)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(rows)
	result := map[documentKey]int{}

	for rows.Next() {
		var collection, path string
		var identifier int
		errors.PanicOnError(rows.Scan(&collection, &path, &identifier))
		result[documentKey{collection, path}] = identifier
	}

	errors.PanicOnError(rows.Err())

	return result
}
