package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunCollisionAliasFix(
	patterns []string,
	directory string,
	ours []string,
	diff bool,
	r *output.Results,
) {
	w := workspace.New(diff, directory)
	aliasThrough(
		patterns,
		directory,
		newAliasMode(false, appliesWhole, ours),
		w,
		r,
	)
	finish(w, r)
}
