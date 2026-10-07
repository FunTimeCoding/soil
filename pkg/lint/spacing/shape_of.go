package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
	"strings"
)

func (s *Spacing) shapeOf(
	line string,
	number int,
) *spacing_shape.Shape {
	trimmed := strings.TrimSpace(withoutRawStrings(line))
	blank := strings.TrimSpace(line) == ""
	topLevel := len(s.blockStack) == 0
	pastTrimmed := strings.TrimSpace(withoutRawStrings(s.pastLine))
	h := spacing_shape.New()
	h.Line = line
	h.Number = number
	h.Trimmed = trimmed
	h.Blank = blank
	h.TopLevel = topLevel
	h.ControlStart = strings.HasPrefix(trimmed, "if ") ||
		strings.HasPrefix(trimmed, "for ") ||
		strings.HasPrefix(trimmed, "switch ") ||
		strings.HasPrefix(trimmed, "select ")
	h.Exit = trimmed == "return" || strings.HasPrefix(trimmed, "return ") ||
		trimmed == "break" || strings.HasPrefix(trimmed, "break ") ||
		trimmed == "continue" || strings.HasPrefix(trimmed, "continue ")
	h.Deferral = strings.HasPrefix(trimmed, "defer ")
	h.TopLevelDeclaration = topLevel && !blank &&
		(strings.HasPrefix(trimmed, "func ") ||
			strings.HasPrefix(trimmed, "type "))
	h.Variable = topLevel && !blank && strings.HasPrefix(trimmed, "var ")
	h.Constant = topLevel && !blank && strings.HasPrefix(trimmed, "const ")
	h.ClosingBrace = strings.HasPrefix(trimmed, "}") ||
		strings.HasPrefix(trimmed, "case ") ||
		trimmed == "default:"
	h.ElseContinuation = strings.HasPrefix(trimmed, "} else")
	h.EndsWithBrace = strings.HasSuffix(trimmed, "{")
	h.PastTrimmed = pastTrimmed
	h.PastOpensBlock = strings.HasSuffix(pastTrimmed, "{") ||
		strings.HasPrefix(pastTrimmed, "case ") ||
		pastTrimmed == "default:" ||
		strings.HasPrefix(pastTrimmed, constant.CommentPrefix)

	return h
}
