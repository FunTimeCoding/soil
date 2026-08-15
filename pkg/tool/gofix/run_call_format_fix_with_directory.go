package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func RunCallFormatFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	RunFormatFixWithDirectory(patterns, directory, false, r)
}
