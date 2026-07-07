package pointer

import (
	"fmt"
	"slices"
	"strings"
)

func Classify(
	s string,
	roots []string,
) string {
	if strings.Contains(s, "://") {
		return Locator
	}

	trimmed, plugin := strings.CutPrefix(
		s,
		fmt.Sprintf("%s/", PluginRootVariable),
	)

	if strings.ContainsAny(trimmed, "<>*$") {
		return Placeholder
	}

	if plugin {
		return Repository
	}

	if strings.HasPrefix(trimmed, "/") {
		if strings.HasPrefix(trimmed, UserPrefix) ||
			strings.HasPrefix(trimmed, HomePrefix) {
			return Absolute
		}

		return Unknown
	}

	if strings.HasPrefix(trimmed, "..") {
		return Sibling
	}

	root, _, _ := strings.Cut(strings.TrimPrefix(trimmed, "./"), "/")

	if !slices.Contains(roots, root) {
		return Unknown
	}

	if IsSymbol(trimmed) {
		return Unknown
	}

	return Repository
}
