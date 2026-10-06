package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func parameterThrough(
	patterns []string,
	directory string,
	w *workspace.Workspace,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := loadThrough(directory, patterns, w)
	edits := findSingleParameterEdits(all, r, w)

	if len(edits) == 0 {
		return
	}

	editThrough(fileSet, edits, directory, w)
}
