package virtual_file_system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func From(directory string) *System {
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

				content, e := os.ReadFile(path)
				errors.PanicOnError(e)
				relative := strings.TrimPrefix(
					path,
					fmt.Sprintf(
						"%s%s",
						directory,
						string(filepath.Separator),
					),
				)
				result.files[relative] = &File{
					Content: content,
					Size:    i.Size(),
					ModTime: i.ModTime(),
					Loaded:  true,
				}

				return nil
			},
		),
	)

	return result
}
