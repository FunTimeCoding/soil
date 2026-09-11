package system

import (
	"os"
	"path/filepath"
)

func FindDirectoryUp(
	currentDirectory string,
	nameToFind string,
) string {
	for {
		if _, e := os.Stat(
			filepath.Join(currentDirectory, nameToFind),
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
