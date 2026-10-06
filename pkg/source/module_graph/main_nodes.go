package module_graph

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"go/build"
	"io/fs"
	"path"
	"path/filepath"
)

func mainNodes(
	c *build.Context,
	root string,
	module string,
) (map[string]*Node, map[string]*Node) {
	result := make(map[string]*Node)
	units := make(map[string]*Node)
	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				directory string,
				d fs.DirEntry,
				e error,
			) error {
				if e != nil || !d.IsDir() {
					return nil
				}

				if Excluded(root, directory) {
					return filepath.SkipDir
				}

				relative, f := filepath.Rel(root, directory)

				if f != nil {
					return nil
				}

				p := module

				if relative != constant.CurrentDirectory {
					p = path.Join(module, filepath.ToSlash(relative))
				}

				n, directoryUnits := importDirectory(c, directory, p)

				if n != nil {
					result[p] = n
				}

				for _, u := range directoryUnits {
					units[u.Path] = u
				}

				return nil
			},
		),
	)

	return result, units
}
