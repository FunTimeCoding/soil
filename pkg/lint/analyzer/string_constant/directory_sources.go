package string_constant

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"os"
	"path/filepath"
	"strings"
)

func DirectorySources(directory string) []string {
	entries, e := os.ReadDir(directory)

	if e != nil {
		return nil
	}

	var result []string

	for _, n := range entries {
		if n.IsDir() || !strings.HasSuffix(n.Name(), soil.GoExtension) {
			continue
		}

		if strings.HasSuffix(n.Name(), soil.TestSuffix) {
			continue
		}

		result = append(result, filepath.Join(directory, n.Name()))
	}

	return result
}
