package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path"
	"slices"
	"strings"
)

func IsBareName(span string) bool {
	name, _, _ := strings.Cut(span, constant.FragmentSeparator)

	return name != "" &&
		!strings.ContainsAny(name, constant.HorizontalWhitespace) &&
		!strings.ContainsAny(name, "/<>*$") &&
		slices.Contains(constant.BareNameExtensions, path.Ext(name))
}
