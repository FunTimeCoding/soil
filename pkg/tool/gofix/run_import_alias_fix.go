package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunImportAliasFix(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	aliasThrough(patterns, directory, newAliasMode(true, nil, nil), w, r)
	finish(w, r)
}
