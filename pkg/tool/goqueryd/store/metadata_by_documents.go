package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Store) metadataByDocuments(
	identifiers []int,
) map[int]map[string]string {
	if len(identifiers) == 0 {
		return nil
	}

	rows, e := s.database.Query(
		fmt.Sprintf(
			`SELECT document_identifier, key, value
			FROM document_metadata
			WHERE document_identifier IN (%s)`,
			placeholders(len(identifiers)),
		),
		intsToAny(identifiers)...,
	)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(rows)
	result := map[int]map[string]string{}

	for rows.Next() {
		var identifier int
		var key, value string
		errors.PanicOnError(rows.Scan(&identifier, &key, &value))

		if result[identifier] == nil {
			result[identifier] = map[string]string{}
		}

		result[identifier][key] = value
	}

	errors.PanicOnError(rows.Err())

	return result
}
