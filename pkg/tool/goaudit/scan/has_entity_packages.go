package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func hasEntityPackages(
	v *virtual_file_system.System,
	path string,
) bool {
	for _, n := range v.MustReadDirectory(path) {
		if n == "basic" ||
			n == "constant" ||
			n == "example" ||
			n == "check" ||
			n == "response" {
			continue
		}

		c := filepath.Join(path, n)

		if !v.DirectoryExists(c) {
			continue
		}

		if v.Has(filepath.Join(c, fmt.Sprintf("%s.go", n))) {
			return true
		}
	}

	return false
}
