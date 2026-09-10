package sweep

import (
	"os"
	"path/filepath"
)

func collectSources() []string {
	base := sourcePath()
	entries, e := os.ReadDir(base)

	if e != nil {
		return nil
	}

	var result []string

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		files, f := filepath.Glob(filepath.Join(base, entry.Name(), "*.jsonl"))

		if f != nil {
			continue
		}

		result = append(result, files...)
	}

	return result
}
