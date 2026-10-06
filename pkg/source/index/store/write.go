package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"os"
	"path/filepath"
)

func (s *Store) Write(
	kind string,
	key string,
	value any,
) {
	p := s.path(kind, key)
	directory := filepath.Dir(p)
	errors.PanicOnError(os.MkdirAll(directory, 0755))
	f, e := os.CreateTemp(directory, "write-*")
	errors.PanicOnError(e)
	_, e = f.Write(notation.Marshal(value))
	errors.PanicOnError(e)
	errors.PanicOnError(f.Close())
	errors.PanicOnError(os.Rename(f.Name(), p))
}
