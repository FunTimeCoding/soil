package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func RunCallFormatFix(
	patterns []string,
	directory string,
	r *output.Results,
) {
	RunFormatFix(patterns, directory, false, r)
}
