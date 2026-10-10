package goreplace

import (
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/match"
	"io"
	"strings"
)

func printDifference(
	out io.Writer,
	content string,
	matches []*match.Match,
) {
	for _, m := range matches {
		write(out, constant.BlockHeader, m.Block.Number, m.Line)
		removed := content[m.Offset : m.Offset+m.Length]
		added := m.Replace

		if m.Block.Opening != constant.SearchMarker {
			removed = strings.TrimSuffix(removed, "\n")
			added = strings.TrimPrefix(strings.TrimSuffix(added, "\n"), "\n")
		}

		if m.Length > 0 {
			for _, line := range strings.Split(removed, "\n") {
				write(out, "%s%s", constant.RemovedPrefix, line)
			}
		}

		if added == "" {
			continue
		}

		for _, line := range strings.Split(added, "\n") {
			write(out, "%s%s", constant.AddedPrefix, line)
		}
	}
}
