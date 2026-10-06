package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
)

func (s *Store) ListSourceTypes() []record.SourceTypeTag {
	rows, e := s.database.Query(
		"SELECT collection, path_prefix, source_type FROM source_type_tag ORDER BY collection, path_prefix",
	)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(rows)
	var result []record.SourceTypeTag

	for rows.Next() {
		var t record.SourceTypeTag
		errors.PanicOnError(
			rows.Scan(&t.Collection, &t.PathPrefix, &t.SourceType),
		)
		result = append(result, t)
	}

	return result
}
