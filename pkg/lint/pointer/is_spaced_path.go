package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path"
	"slices"
	"strings"
)

func isSpacedPath(s string) bool {
	first := strings.Fields(s)

	return len(first) > 0 &&
		strings.Contains(first[0], "/") &&
		slices.Contains(constant.BareNameExtensions, path.Ext(s))
}
