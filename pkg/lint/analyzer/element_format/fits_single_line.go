package element_format

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/token"
)

func FitsSingleLine(
	fileSet *token.FileSet,
	e *Elements,
	source []byte,
) bool {
	for _, item := range e.Items {
		if fileSet.Position(item.Pos()).Line != fileSet.Position(item.End()).Line {
			return false
		}
	}

	p := fileSet.Position(e.Position)
	lineStart := p.Offset - (p.Column - 1)
	width := 0

	for i := lineStart; i < p.Offset; i++ {
		if source[i] == '\t' {
			width += constant.TabWidth
		} else {
			width++
		}
	}

	width += fileSet.Position(e.Open).Offset - p.Offset
	width++

	for i, item := range e.Items {
		start := fileSet.Position(item.Pos()).Offset
		end := fileSet.Position(item.End()).Offset
		width += end - start

		if i < len(e.Items)-1 {
			width += 2
		}
	}

	if e.Ellipsis {
		width += 3
	}

	width += e.Padding
	width++
	closeOffset := fileSet.Position(e.Close).Offset

	for i := closeOffset + 1; i < len(source) && source[i] != '\n'; i++ {
		width++
	}

	return width <= constant.MaxLineLength
}
