package measure

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
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
		if strings.Contains(s, stringsConstant.Slash) {
			if relative == strings.TrimSuffix(s, stringsConstant.Slash) ||
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
