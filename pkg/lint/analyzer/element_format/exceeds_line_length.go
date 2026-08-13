package element_format

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/token"
)

func exceedsLineLength(
	fileSet *token.FileSet,
	e *Elements,
	source []byte,
) bool {
	p := fileSet.Position(e.Close)
	lineStart := p.Offset - (p.Column - 1)
	width := 0

	for i := lineStart; i < p.Offset; i++ {
		if source[i] == '\t' {
			width += constant.TabWidth
		} else {
			width++
		}
	}

	width++

	return width > constant.MaxLineLength
}
