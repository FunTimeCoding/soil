package watcher

import (
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func sourceDirectories() []string {
	base := filepath.Join(system.Home(), ".claude", "projects")
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
