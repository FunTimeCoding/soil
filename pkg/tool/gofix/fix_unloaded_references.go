package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func FixUnloadedReferences(
	violations []Violation,
	loadedFiles map[string]bool,
	directory string,
	r *output.Results,
) {
	w := workspace.New(false, directory)
	referencesThrough(violations, loadedFiles, directory, r, w)
	finish(w, r)
}
