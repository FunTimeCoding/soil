package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func (s *Store) path(
	kind string,
	key string,
) string {
	return filepath.Join(s.directory, kind, key[:2], join.Empty(key, ".json"))
}
