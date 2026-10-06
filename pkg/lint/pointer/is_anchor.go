package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func isAnchor(target string) bool {
	return len(target) > len(constant.FragmentSeparator) &&
		strings.HasPrefix(target, constant.FragmentSeparator) &&
		!strings.ContainsAny(target, constant.HorizontalWhitespace)
}
