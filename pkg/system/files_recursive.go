package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func FilesRecursive(
	root string,
	verbose bool,
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

				if verbose {
					fmt.Printf("FilesRecursive walk: %s\n", path)
				}

				if !i.IsDir() {
					result = append(result, path)
				}

				return nil
			},
		),
	)

	return result
}
