package snapshot

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/types/snapshot_stamp"
	"io/fs"
	"path/filepath"
	"strings"
)

func walk(root string) map[string]snapshot_stamp.Stamp {
	result := make(map[string]snapshot_stamp.Stamp)
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

				if !strings.HasSuffix(path, soil.GoExtension) &&
					path != filepath.Join(root, go_modConstant.ModFile) &&
					path != filepath.Join(root, go_modConstant.SumFile) {
					return nil
				}

				i, g := d.Info()

				if g != nil {
					return nil
				}

				result[path] = snapshot_stamp.Stamp{
					Size:     i.Size(),
					Modified: i.ModTime(),
				}

				return nil
			},
		),
	)

	return result
}
