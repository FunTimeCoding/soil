package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"strings"
)

func isIgnoredPath(path string) bool {
	segments := strings.Split(path, "/")

	for _, s := range segments[:len(segments)-1] {
		if strings.HasPrefix(s, constant.IgnoredPrefix) {
			return true
		}
	}

	return false
}
