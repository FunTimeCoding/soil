package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunConstantAliasFix(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	aliasThrough(
		patterns,
		directory,
		newAliasMode(false, defaultApplies, nil),
		w,
		r,
	)
	finish(w, r)
}
