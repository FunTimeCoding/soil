package repository

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func siblingContents(directory string) []string {
	var result []string
	errors.PanicOnError(
		filepath.Walk(
			directory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if i.IsDir() ||
					i.Mode()&os.ModeSymlink != 0 ||
					strings.HasSuffix(path, constant.MarkdownExtension) {
					return nil
				}

				b, f := os.ReadFile(path)
				errors.PanicOnError(f)
				result = append(result, string(b))

				return nil
			},
		),
	)

	return result
}
