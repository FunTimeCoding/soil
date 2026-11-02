package sweep

import (
	"os"
	"path/filepath"
)

func SourceDirectories() []string {
	base := sourcePath()
	entries, e := os.ReadDir(base)

	if e != nil {
		return nil
	}

	var result []string

	for _, entry := range entries {
		if entry.IsDir() {
			result = append(result, filepath.Join(base, entry.Name()))
		}
	}

	return result
}
