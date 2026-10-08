package goreplace

import (
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/match"
	"io"
	"strings"
)

func printDifference(
	out io.Writer,
	matches []*match.Match,
) {
	for _, m := range matches {
		write(out, constant.BlockHeader, m.Block.Number, m.Line)

		for _, line := range strings.Split(m.Block.Search, "\n") {
			write(out, "%s%s", constant.RemovedPrefix, line)
		}

		if m.Block.Replace == "" {
			continue
		}

		for _, line := range strings.Split(m.Block.Replace, "\n") {
			write(out, "%s%s", constant.AddedPrefix, line)
		}
	}
}
