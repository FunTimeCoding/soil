package repository

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
	"path/filepath"
	"strings"
)

func siblingTests(
	directory string,
	ignored func(string) bool,
) []string {
	var result []string
	errors.PanicOnError(
		filepath.WalkDir(
			directory,
			func(
				path string,
				d fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				relative, f := filepath.Rel(directory, path)

				if f != nil {
					return f
				}

				if d.IsDir() {
					if relative != soil.CurrentDirectory &&
						(strings.HasPrefix(d.Name(), soil.CurrentDirectory) ||
							ignored != nil && ignored(relative)) {
						return filepath.SkipDir
					}

					return nil
				}

				if strings.HasSuffix(path, soil.TestSuffix) {
					result = append(result, path)
				}

				return nil
			},
		),
	)

	return result
}
