package reach

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"os"
	"strings"
)

func hasGoFiles(directory string) bool {
	entries, e := os.ReadDir(directory)

	if e != nil {
		return false
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), soil.GoExtension) {
			return true
		}
	}

	return false
}
