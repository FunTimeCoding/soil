package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (s *Spacing) shapeOf(
	line string,
	number int,
) *shape {
	trimmed := strings.TrimSpace(withoutRawStrings(line))
	blank := strings.TrimSpace(line) == ""
	topLevel := len(s.blockStack) == 0
	pastTrimmed := strings.TrimSpace(withoutRawStrings(s.pastLine))

	return &shape{
		line:     line,
		number:   number,
		trimmed:  trimmed,
		blank:    blank,
		topLevel: topLevel,
		controlStart: strings.HasPrefix(trimmed, "if ") ||
			strings.HasPrefix(trimmed, "for ") ||
			strings.HasPrefix(trimmed, "switch ") ||
			strings.HasPrefix(trimmed, "select "),
		exit: trimmed == "return" || strings.HasPrefix(
			trimmed,
			"return ",
		) ||
			trimmed == "break" || strings.HasPrefix(trimmed, "break ") ||
			trimmed == "continue" || strings.HasPrefix(trimmed, "continue "),
		deferral: strings.HasPrefix(trimmed, "defer "),
		topLevelDeclaration: topLevel && !blank && (strings.HasPrefix(
			trimmed,
			"func ",
		) ||
			strings.HasPrefix(trimmed, "type ")),
		variable: topLevel && !blank && strings.HasPrefix(trimmed, "var "),
		constant: topLevel &&
			!blank &&
			strings.HasPrefix(trimmed, "const "),
		closingBrace: strings.HasPrefix(trimmed, "}") ||
			strings.HasPrefix(trimmed, "case ") ||
			trimmed == "default:",
		elseContinuation: strings.HasPrefix(trimmed, "} else"),
		endsWithBrace:    strings.HasSuffix(trimmed, "{"),
		pastTrimmed:      pastTrimmed,
		pastOpensBlock: strings.HasSuffix(pastTrimmed, "{") ||
			strings.HasPrefix(pastTrimmed, "case ") ||
			pastTrimmed == "default:" ||
			strings.HasPrefix(pastTrimmed, constant.CommentPrefix),
	}
}
