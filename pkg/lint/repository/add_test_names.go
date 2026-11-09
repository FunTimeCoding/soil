package repository

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"os"
	"path/filepath"
)

func addTestNames(
	names map[string]map[string]bool,
	path string,
) {
	b, e := os.ReadFile(path)

	if e != nil {
		return
	}

	directory := filepath.Dir(filepath.Clean(path))

	for _, m := range constant.TestDeclaration.FindAllSubmatch(b, -1) {
		name := string(m[1])

		if names[name] == nil {
			names[name] = map[string]bool{}
		}

		names[name][directory] = true
	}
}
