package element_format

import (
	"go/token"
	"strings"
)

func Indent(
	fileSet *token.FileSet,
	e *Elements,
	source []byte,
) string {
	p := fileSet.Position(e.Position)
	lineStart := p.Offset - (p.Column - 1)
	count := 0

	for i := lineStart; i < len(source) && source[i] == '\t'; i++ {
		count++
	}

	return strings.Repeat("\t", count)
}
