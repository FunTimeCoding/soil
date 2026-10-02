package goflow

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func rewrap(
	source []byte,
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
	var text []string

	for _, line := range strings.Split(string(source[start:last.Stop]), "\n") {
		longest = max(longest, len(line))
	}

	for _, s := range segments {
		text = append(text, strings.TrimSpace(string(source[s.Start:s.Stop])))
	}

	if longest <= width {
		return nil
	}

	wrapped := wrap(strings.Join(text, " "), width-len(prefix))

	for i := 1; i < len(wrapped); i++ {
		wrapped[i] = join.Empty(strings.Repeat(" ", len(prefix)), wrapped[i])
	}

	wrapped[0] = join.Empty(prefix, wrapped[0])

	return &patch{
		start:   start,
		stop:    last.Stop,
		content: join.NewLine(wrapped),
	}
}
