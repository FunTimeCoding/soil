package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunSingleParameterFix(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	parameterThrough(patterns, directory, w, r)
	finish(w, r)
}
