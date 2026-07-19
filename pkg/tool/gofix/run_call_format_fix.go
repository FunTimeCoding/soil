package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runCallFormatFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	RunCallFormatFixWithDirectory(patterns, "", r)
}
