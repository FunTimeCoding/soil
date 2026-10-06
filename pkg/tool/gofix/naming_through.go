package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func namingThrough(
	patterns []string,
	directory string,
	w *workspace.Workspace,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := loadThrough(directory, patterns, w)
	edits := findVariableNamingEdits(fileSet, all, r)

	if len(edits) == 0 {
		return
	}

	editThrough(fileSet, edits, directory, w)
}
