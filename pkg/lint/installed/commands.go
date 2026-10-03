package installed

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"os"
	"path/filepath"
)

func commands(modules map[string]string) map[string][]string {
	result := make(map[string][]string)

	for module, directory := range modules {
		entries, e := os.ReadDir(
			filepath.Join(directory, constant.CommandDirectory),
		)

		if e != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				result[entry.Name()] = append(result[entry.Name()], module)
			}
		}
	}

	return result
}
