package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunConstantAliasFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	aliasThrough(
		patterns,
		directory,
		newAliasMode(false, appliesByDefault, nil),
		w,
		r,
	)
	finish(w, r)
}
