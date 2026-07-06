package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"strings"
)

func Skipped(
	o *option.Lint,
	path string,
) bool {
	if o.Count == 0 {
		return false
	}

	for _, p := range o.Skips {
		if strings.Contains(p, separator.Dot) &&
			!strings.Contains(p, separator.Slash) {
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
