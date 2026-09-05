package rename_test_homes

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
	"path/filepath"
	"strings"
)

func testHomeDirectories(
	root string,
	name string,
) []string {
	var result []string
	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				current string,
				entry fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				if !entry.IsDir() {
					return nil
				}

				if strings.HasPrefix(
					entry.Name(),
					constant.CurrentDirectory,
				) || entry.Name() == "testdata" {
					return filepath.SkipDir
				}

				if entry.Name() == name {
					relative, f := filepath.Rel(root, current)

					if f != nil {
						return f
					}

					result = append(result, relative)

					return filepath.SkipDir
				}

				return nil
			},
		),
	)

	return result
}
