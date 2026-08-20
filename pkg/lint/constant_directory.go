package lint

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"path/filepath"
	"strings"
)

// A file under a constant/ role directory holds constants across
// concept files - the package is named after the domain, not
// constant, so the package-name exemption does not apply.
func constantDirectory(path string) bool {
	for s := range strings.SplitSeq(
		filepath.ToSlash(filepath.Dir(path)),
		constant.Slash,
	) {
		if s == "constant" {
			return true
		}
	}

	return false
}
