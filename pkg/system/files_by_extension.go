package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"slices"
)

func FilesByExtension(
	root string,
	extension ...string,
) []string {
	var result []string
	errors.PanicOnError(
		filepath.Walk(
			root,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if !i.IsDir() && slices.Contains(
					extension,
					filepath.Ext(path),
				) {
					result = append(result, path)
				}

				return nil
			},
		),
	)

	return result
}
