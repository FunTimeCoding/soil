package installed

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"os"
	"path/filepath"
)

func PathBinaries(known map[string]bool) []*Binary {
	seen := make(map[string]bool)
	var result []*Binary

	for _, directory := range filepath.SplitList(environment.Optional("PATH")) {
		entries, e := os.ReadDir(directory)

		if e != nil {
			continue
		}

		for _, entry := range entries {
			name := entry.Name()

			if entry.IsDir() || seen[name] {
				continue
			}

			seen[name] = true

			if !known[name] {
				continue
			}

			if b, okay := Read(filepath.Join(directory, name)); okay {
				result = append(result, b)
			}
		}
	}

	return result
}
