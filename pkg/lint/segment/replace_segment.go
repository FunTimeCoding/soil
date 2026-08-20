package segment

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
	"unicode"
)

func ReplaceSegment(name, old, replacement string) string {
	spans := segmentSpans(name)
	var target *segmentSpan

	for i := range spans {
		if spans[i].lower == old {
			target = &spans[i]

			break
		}
	}

	if target == nil {
		return name
	}

	firstUpper := unicode.IsUpper(rune(name[target.start]))
	words := split.Underscore(replacement)
	underscore := strings.Contains(name, constant.Underscore)
	var b strings.Builder

	for i, w := range words {
		if i > 0 && underscore {
			b.WriteByte('_')
		}

		if i == 0 && firstUpper {
			b.WriteString(capitalize(w))
		} else if i > 0 && !underscore {
			b.WriteString(capitalize(w))
		} else {
			b.WriteString(w)
		}
	}

	return join.Empty(name[:target.start], b.String(), name[target.end:])
}
