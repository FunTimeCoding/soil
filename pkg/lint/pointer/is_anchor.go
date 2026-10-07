package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func IsAnchor(target string) bool {
	return len(target) > len(constant.FragmentSeparator) &&
		strings.HasPrefix(target, constant.FragmentSeparator) &&
		!strings.ContainsAny(target, constant.HorizontalWhitespace)
}
