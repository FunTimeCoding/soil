package system

import (
	"github.com/funtimecoding/soil/pkg/system/join"
	"os"
	"path/filepath"
)

func FindDirectoryUp(
	currentDirectory string,
	nameToFind string,
) string {
	for {
		if _, e := os.Stat(
			join.Absolute(currentDirectory, nameToFind),
		); e == nil {
			return currentDirectory
		}

		parent := filepath.Dir(currentDirectory)

		if parent == currentDirectory {
			return ""
		}

		currentDirectory = parent
	}
}
