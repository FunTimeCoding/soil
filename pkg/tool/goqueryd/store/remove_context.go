package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func (s *Store) RemoveContext(
	collection string,
	pathPrefix string,
) bool {
	if !strings.HasPrefix(pathPrefix, separator.Slash) {
		pathPrefix = join.Empty(separator.Slash, pathPrefix)
	}

	if !strings.HasSuffix(pathPrefix, separator.Slash) {
		pathPrefix = join.Empty(pathPrefix, separator.Slash)
	}

	result, e := s.database.Exec(
		"DELETE FROM context WHERE collection = ? AND path_prefix = ?",
		collection,
		pathPrefix,
	)
	errors.PanicOnError(e)
	affected, e := result.RowsAffected()
	errors.PanicOnError(e)

	return affected > 0
}
