package gofix

import (
	"github.com/funtimecoding/soil/pkg/console"
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
	lines := diffLines(original, modified)
	shown := make([]bool, len(lines))

	for i, l := range lines {
		if l.Mark == " " {
			continue
		}

		for j := i - constant.DiffContext; j <= i+constant.DiffContext; j++ {
			if j >= 0 && j < len(lines) {
				shown[j] = true
			}
		}
	}

	gap := false

	for i, l := range lines {
		if !shown[i] {
			gap = true

			continue
		}

		if gap {
			console.Format("@@\n")
			gap = false
		}

		console.Format("%s%s\n", l.Mark, l.Text)
	}
}
