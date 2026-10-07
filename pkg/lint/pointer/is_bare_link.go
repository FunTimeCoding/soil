package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func IsBareLink(target string) bool {
	return target != "" &&
		!strings.ContainsAny(target, constant.HorizontalWhitespace) &&
		!strings.Contains(target, constant.LocatorSeparator) &&
		!strings.HasPrefix(target, constant.FragmentSeparator) &&
		!strings.ContainsAny(target, "<>*$") &&
		strings.ContainsRune(target, '.')
}
