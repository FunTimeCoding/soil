package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runCallFormatFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := load(directory, patterns)
	edits := findCallFormatEdits(all, r)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, false)
}
