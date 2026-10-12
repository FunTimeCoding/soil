package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func aliasThrough(
	patterns []string,
	directory string,
	realias bool,
	w *workspace.Workspace,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := loadThrough(directory, patterns, w)
	edits, handled := findImportAliasSuggestions(fileSet, all, realias, r)
	edits = append(edits, findImportAliasEdits(fileSet, all, handled, r)...)

	if len(edits) == 0 {
		return
	}

	editThrough(fileSet, edits, directory, w)
}
