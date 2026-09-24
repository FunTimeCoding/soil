package measure

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path"
	"strings"
)

func skipped(
	skips []string,
	relative string,
) bool {
	name := path.Base(relative)

	for _, s := range skips {
		if strings.Contains(s, constant.Slash) {
			if relative == strings.TrimSuffix(s, constant.Slash) ||
				strings.HasPrefix(relative, s) {
				return true
			}

			continue
		}

		if system.Match(s, name) {
			return true
		}
	}

	return false
}
