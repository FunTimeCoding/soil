package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunImportAliasFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	aliasThrough(patterns, directory, w, r)
	finish(w, r)
}
