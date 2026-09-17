package virtual_file_system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func FromMetadata(directory string) *System {
	result := New()
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

				if i.IsDir() || i.Mode()&os.ModeSymlink != 0 {
					return nil
				}

				result.AddMetadata(
					strings.TrimPrefix(
						path,
						fmt.Sprintf(
							"%s%s",
							directory,
							string(filepath.Separator),
						),
					),
					i.Size(),
					i.ModTime(),
				)

				return nil
			},
		),
	)

	return result
}
