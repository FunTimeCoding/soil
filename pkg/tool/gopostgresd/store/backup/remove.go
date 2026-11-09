package backup

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"os"
)

func (s *Store) Remove(
	instance string,
	name string,
) error {
	path, e := s.Existing(instance, name)

	if e != nil {
		return e
	}

	if f := os.Remove(path); f != nil {
		return f
	}

	g := os.Remove(join.Empty(path, constant.ChecksumExtension))

	if g != nil && !os.IsNotExist(g) {
		return g
	}

	return nil
}
