package snapshot

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	goModule "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"io/fs"
	"path/filepath"
	"strings"
)

func walk(root string) map[string]stamp {
	result := make(map[string]stamp)
	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				path string,
				d fs.DirEntry,
				f error,
			) error {
				if f != nil {
					return nil
				}

				if d.IsDir() {
					if module_graph.Excluded(root, path) {
						return filepath.SkipDir
					}

					return nil
				}

				if !strings.HasSuffix(path, constant.GoExtension) &&
					path != filepath.Join(root, goModule.ModFile) &&
					path != filepath.Join(root, goModule.SumFile) {
					return nil
				}

				i, g := d.Info()

				if g != nil {
					return nil
				}

				result[path] = stamp{size: i.Size(), modified: i.ModTime()}

				return nil
			},
		),
	)

	return result
}
