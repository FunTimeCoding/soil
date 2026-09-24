package measure

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
	"io/fs"
	"path/filepath"
)

func walk(
	root string,
	skips []string,
) []string {
	var out []string
	ignored := git.IgnoreMatcher(root)
	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				path string,
				d fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				if path == root {
					return nil
				}

				relative := filepath.ToSlash(system.RelativePath(root, path))

				if d.IsDir() {
					if skipped(skips, relative) || ignored(relative) {
						return filepath.SkipDir
					}

					return nil
				}

				if d.Type()&fs.ModeSymlink != 0 ||
					skipped(skips, relative) ||
					ignored(relative) {
					return nil
				}

				out = append(out, path)

				return nil
			},
		),
	)

	return out
}
