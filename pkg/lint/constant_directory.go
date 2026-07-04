package lint

import (
	"path/filepath"
	"strings"
)

// A file under a constant/ role directory holds constants across
// concept files - the package is named after the domain, not
// constant, so the package-name exemption does not apply.
func constantDirectory(path string) bool {
	for _, s := range strings.Split(
		filepath.ToSlash(filepath.Dir(path)),
		"/",
	) {
		if s == "constant" {
			return true
		}
	}

	return false
}
