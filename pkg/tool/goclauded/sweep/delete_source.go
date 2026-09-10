package sweep

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func DeleteSource(sessionIdentifier string) []string {
	var result []string
	base := sourcePath()
	entries, e := os.ReadDir(base)

	if e != nil {
		return result
	}

	target := join.Empty(sessionIdentifier, constant.NotationLogExtension)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		path := filepath.Join(base, entry.Name(), target)

		if f := os.Remove(path); f == nil {
			result = append(result, path)
		}

		directory := filepath.Join(base, entry.Name(), sessionIdentifier)

		if _, f := os.Stat(directory); f != nil {
			continue
		}

		if f := os.RemoveAll(directory); f == nil {
			result = append(result, directory)
		}
	}

	return result
}
