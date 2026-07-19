package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func RunCallFormatFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := Load(directory, patterns)
	edits := findCallFormatEdits(all, r)

	if len(edits) == 0 {
		return
	}

	ApplyEdits(fileSet, edits, directory, false)
}
