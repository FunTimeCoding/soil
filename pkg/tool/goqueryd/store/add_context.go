package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func (s *Store) AddContext(
	collection string,
	pathPrefix string,
	description string,
) {
	if !strings.HasPrefix(pathPrefix, separator.Slash) {
		pathPrefix = join.Empty(separator.Slash, pathPrefix)
	}

	if !strings.HasSuffix(pathPrefix, separator.Slash) {
		pathPrefix = join.Empty(pathPrefix, separator.Slash)
	}

	_, e := s.database.Exec(
		`INSERT INTO context (collection, path_prefix, description)
		VALUES (?, ?, ?)
		ON CONFLICT(collection, path_prefix) DO UPDATE SET
			description = excluded.description`,
		collection,
		pathPrefix,
		description,
	)
	errors.PanicOnError(e)
}
