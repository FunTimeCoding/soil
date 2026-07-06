package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) SetSourceType(
	collection string,
	pathPrefix string,
	sourceType string,
) {
	pathPrefix = normalizePathPrefix(pathPrefix)

	if sourceType == "" {
		_, e := s.database.Exec(
			"DELETE FROM source_type_tag WHERE collection = ? AND path_prefix = ?",
			collection,
			pathPrefix,
		)
		errors.PanicOnError(e)

		return
	}

	_, e := s.database.Exec(
		`INSERT INTO source_type_tag (collection, path_prefix, source_type)
		VALUES (?, ?, ?)
		ON CONFLICT(collection, path_prefix) DO UPDATE SET source_type = excluded.source_type`,
		collection,
		pathPrefix,
		sourceType,
	)
	errors.PanicOnError(e)
}
