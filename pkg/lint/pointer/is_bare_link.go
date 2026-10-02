package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func isBareLink(target string) bool {
	return target != "" &&
		!strings.ContainsAny(target, constant.HorizontalWhitespace) &&
		!strings.Contains(target, "://") &&
		!strings.HasPrefix(target, "#") &&
		!strings.ContainsAny(target, "<>*$") &&
		strings.ContainsRune(target, '.')
}
