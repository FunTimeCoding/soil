package warm_tree

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/source/types/tree_entry"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
)

func walk(root string) (
	map[string]*tree_entry.Entry,
	map[string][]string,
) {
	entries := make(map[string]*tree_entry.Entry)
	children := make(map[string][]string)
	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				p string,
				d fs.DirEntry,
				e error,
			) error {
				if e != nil {
					if p == root {
						return e
					}

					return nil
				}

				relative, f := filepath.Rel(root, p)
				errors.PanicOnError(f)
				relative = filepath.ToSlash(relative)

				if d.IsDir() && relative != soil.CurrentDirectory &&
					d.Name() == gitConstant.Directory {
					return filepath.SkipDir
				}

				linked := d.Type()&fs.ModeSymlink != 0
				var i fs.FileInfo
				var g error

				if linked {
					i, g = os.Stat(p)
				} else {
					i, g = d.Info()
				}

				if g != nil || linked && !i.IsDir() {
					return nil
				}

				entries[relative] = tree_entry.New(
					i.Size(),
					i.ModTime(),
					i.IsDir(),
					linked,
				)

				if relative != soil.CurrentDirectory {
					parent := path.Dir(relative)
					children[parent] = append(children[parent], d.Name())
				}

				return nil
			},
		),
	)

	for _, names := range children {
		slices.Sort(names)
	}

	return entries, children
}
