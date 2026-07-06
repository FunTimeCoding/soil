package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runSingleParameterFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := load(directory, patterns)
	edits := findSingleParameterEdits(all, r)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, false)
}
