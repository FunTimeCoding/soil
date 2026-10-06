package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
)

func (s *Store) documentIdentifiers(keys []record.DocumentKey) map[record.DocumentKey]int {
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
	result := map[record.DocumentKey]int{}

	for rows.Next() {
		var collection, path string
		var identifier int
		errors.PanicOnError(rows.Scan(&collection, &path, &identifier))
		result[record.DocumentKey{Collection: collection, Path: path}] = identifier
	}

	errors.PanicOnError(rows.Err())

	return result
}
