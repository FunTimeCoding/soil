package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runImportAliasFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := load(directory, patterns)
	edits := findImportAliasEdits(fileSet, all, r)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, diff)
}
