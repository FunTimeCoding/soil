package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"os"
	"path/filepath"
	"strings"
)

func EmptyDirectories(
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
					fmt.Printf("EmptyDirectories walk: %s\n", path)
				}

				if i.IsDir() {
					if !strings.Contains(path, constant.Directory) {
						if IsEmptyDirectory(path) {
							result = append(result, path)
						}
					}
				}

				return nil
			},
		),
	)

	return result
}
