package build_tag

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
	"path/filepath"
	"strings"
)

func scan(directory string) map[string][]string {
	result := make(map[string][]string)
	errors.PanicOnError(
		filepath.WalkDir(
			directory,
			func(
				path string,
				n fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if n.IsDir() {
					if skipped(n.Name()) {
						return filepath.SkipDir
					}

					return nil
				}

				if !strings.HasSuffix(path, constant.GoExtension) {
					return nil
				}

				if tags := Extract(path); len(tags) > 0 {
					result[path] = tags
				}

				return nil
			},
		),
	)

	return result
}
