package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
)

func ApplyEdits(
	fileSet *token.FileSet,
	edits []Edit,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	editThrough(fileSet, edits, directory, w)
	finish(w, r)
}
