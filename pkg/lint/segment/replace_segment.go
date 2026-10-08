package segment

import (
	"github.com/funtimecoding/soil/pkg/lint/types/segment_span"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
	"unicode"
)

func ReplaceSegment(name, old, replacement string) string {
	var target *segment_span.Span

	for _, s := range segmentSpans(name) {
		if s.Lower == old {
			target = s

			break
		}
	}

	if target == nil {
		return name
	}

	firstUpper := unicode.IsUpper(rune(name[target.Start]))
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

	return join.Empty(name[:target.Start], b.String(), name[target.End:])
}
