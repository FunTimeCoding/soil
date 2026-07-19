package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func RunImportAliasFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := Load(directory, patterns)
	edits := findImportAliasEdits(fileSet, all, r)

	if len(edits) == 0 {
		return
	}

	ApplyEdits(fileSet, edits, directory, diff)
}
