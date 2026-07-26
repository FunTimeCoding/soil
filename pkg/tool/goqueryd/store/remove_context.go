package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (s *Store) RemoveContext(
	collection string,
	pathPrefix string,
) bool {
	if !strings.HasPrefix(pathPrefix, constant.Slash) {
		pathPrefix = join.Empty(constant.Slash, pathPrefix)
	}

	if !strings.HasSuffix(pathPrefix, constant.Slash) {
		pathPrefix = join.Empty(pathPrefix, constant.Slash)
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
