package goflow

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
	"unicode/utf8"
)

func rewrap(
	source []byte,
	literal []bool,
	n ast.BlockNode,
	width int,
) *patch {
	segments := n.Source()

	if len(segments) == 0 {
		return nil
	}

	first := segments[0]
	last := segments[len(segments)-1]
	start := lineStart(source, first.Start)
	prefix := string(source[start:first.Start])

	if strings.Contains(prefix, "\t") {
		return nil
	}

	if strings.HasPrefix(string(source[first.Start:first.Stop]), constant.Pipe) {
		return nil
	}

	longest := 0

	for _, line := range strings.Split(string(source[start:last.Stop]), "\n") {
		longest = max(longest, utf8.RuneCountInString(line))
	}

	if longest <= width {
		return nil
	}

	indent := utf8.RuneCountInString(prefix)
	wrapped := wrap(
		bind(units(source, literal, first.Start, last.Stop)),
		width-indent,
	)

	for i := 1; i < len(wrapped); i++ {
		wrapped[i] = join.Empty(strings.Repeat(" ", indent), wrapped[i])
	}

	wrapped[0] = join.Empty(prefix, wrapped[0])

	return &patch{
		start:   start,
		stop:    last.Stop,
		content: join.NewLine(wrapped),
	}
}
