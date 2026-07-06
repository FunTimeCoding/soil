package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runVariableNamingFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runVariableNamingFixWithDirectory(patterns, "", diff, r)
}
