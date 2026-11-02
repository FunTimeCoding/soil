package rename_test_homes

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func hasGoFiles(
	directory string,
	relative string,
) bool {
	entries, e := os.ReadDir(filepath.Join(directory, relative))
	errors.PanicOnError(e)

	for _, entry := range entries {
		if !entry.IsDir() &&
			strings.HasSuffix(entry.Name(), soil.GoExtension) {
			return true
		}
	}

	return false
}
