package state_tree

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"io/fs"
	"path/filepath"
	"strings"
)

func (t *Tree) loadRoot(directory string) {
	errors.PanicOnError(
		filepath.WalkDir(
			directory,
			func(
				path string,
				entry fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				if entry.IsDir() {
					return nil
				}

				relative, f := filepath.Rel(directory, path)

				if f != nil {
					return f
				}

				relative = filepath.ToSlash(relative)
				t.files[relative] = path

				if relative != constant.TopFile &&
					strings.HasSuffix(relative, constant.StateExtension) {
					t.states[stateName(relative)] = relative
				}

				return nil
			},
		),
	)
}
