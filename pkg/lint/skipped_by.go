package lint

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"strings"
)

func SkippedBy(
	skips []string,
	path string,
) bool {
	for _, p := range skips {
		if strings.Contains(p, stringsConstant.Dot) &&
			!strings.Contains(p, stringsConstant.Slash) {
			if system.Match(p, filepath.Base(path)) {
				return true
			}
		} else if strings.HasPrefix(path, p) ||
			strings.Contains(path, fmt.Sprintf("/%s", p)) {
			return true
		}
	}

	return false
}
