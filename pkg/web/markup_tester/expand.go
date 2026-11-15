package markup_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/style_sheet"
	"strings"
)

func expand(s string) []string {
	for _, prefix := range []string{":where(", ":is("} {
		start := strings.Index(s, prefix)

		if start < 0 {
			continue
		}

		open := start + len(prefix) - 1
		last := closing(s, open)
		var result []string

		for _, a := range style_sheet.SplitTop(s[open+1:last], ',') {
			result = append(
				result,
				expand(join.Empty(s[:start], a, s[last+1:]))...)
		}

		return result
	}

	return []string{strings.TrimSpace(s)}
}
