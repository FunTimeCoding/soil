package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunFormatFix(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	formatThrough(patterns, directory, w, r)
	finish(w, r)
}
