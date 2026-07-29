package pointer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"slices"
	"strings"
)

func Classify(
	s string,
	roots []string,
) string {
	if strings.Contains(s, "://") {
		return constant.PointerLocator
	}

	trimmed, plugin := strings.CutPrefix(
		s,
		fmt.Sprintf("%s/", constant.PluginRootVariable),
	)

	if strings.ContainsAny(trimmed, "<>*$") {
		return constant.PointerPlaceholder
	}

	if plugin {
		return constant.PointerRepository
	}

	if strings.HasPrefix(trimmed, "/") {
		if strings.HasPrefix(trimmed, constant.UserPathPrefix) ||
			strings.HasPrefix(trimmed, constant.HomePathPrefix) {
			return constant.PointerAbsolute
		}

		return constant.PointerUnknown
	}

	if strings.HasPrefix(trimmed, "..") {
		return constant.PointerSibling
	}

	root, _, _ := strings.Cut(strings.TrimPrefix(trimmed, "./"), "/")

	if !slices.Contains(roots, root) {
		return constant.PointerUnknown
	}

	if IsSymbol(trimmed) {
		return constant.PointerUnknown
	}

	return constant.PointerRepository
}
