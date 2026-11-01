package gofix

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/text/line_diff"
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
)

func printDiff(
	path string,
	original []byte,
	modified []byte,
) {
	if string(original) == string(modified) {
		return
	}

	console.Format("--- %s\n+++ %s\n", path, path)

	for _, l := range line_diff.Hunks(
		line_diff.Lines(string(original), string(modified)),
		constant.DiffContext,
	) {
		console.Format("%s\n", l)
	}
}
