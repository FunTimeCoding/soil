package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runSingleParameterFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runSingleParameterFixWithDirectory(patterns, "", r)
}
