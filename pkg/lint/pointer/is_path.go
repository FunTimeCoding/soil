package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func IsPath(s string) bool {
	if s == "" {
		return false
	}

	if strings.ContainsAny(s, constant.HorizontalWhitespace) {
		return isSpacedPath(s)
	}

	if strings.Contains(s, "/") {
		return true
	}

	return strings.HasPrefix(s, constant.SchemeGo) &&
		len(s) > len(constant.SchemeGo)
}
