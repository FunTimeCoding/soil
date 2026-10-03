package installed

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func modules(parent string) map[string]string {
	result := make(map[string]string)
	entries, e := os.ReadDir(parent)

	if e != nil {
		return result
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		directory := filepath.Join(parent, entry.Name())
		b, f := os.ReadFile(filepath.Join(directory, constant.ModFile))

		if f != nil {
			continue
		}

		if p := modfile.ModulePath(b); p != "" {
			result[p] = directory
		}
	}

	return result
}
