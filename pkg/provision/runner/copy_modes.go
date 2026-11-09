package runner

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"io/fs"
	"path/filepath"
)

func copyModes(
	source string,
	destination string,
) {
	errors.PanicOnError(
		filepath.WalkDir(
			source,
			func(path string, d fs.DirEntry, e error) error {
				if e != nil {
					return e
				}

				if d.Type()&fs.ModeSymlink != 0 {
					return nil
				}

				i, f := d.Info()

				if f != nil {
					return f
				}

				relative, g := filepath.Rel(source, path)

				if g != nil {
					return g
				}

				system.ChangeMode(
					filepath.Join(destination, relative),
					i.Mode().Perm(),
				)

				return nil
			},
		),
	)
}
