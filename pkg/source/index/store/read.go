package store

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"os"
)

func (s *Store) Read(
	kind string,
	key string,
	target any,
) bool {
	b, e := os.ReadFile(s.path(kind, key))

	if e != nil {
		return false
	}

	return notation.DecodeBytes(b, target) == nil
}
