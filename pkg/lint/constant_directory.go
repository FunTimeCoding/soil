package lint

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"path/filepath"
	"strings"
)

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
