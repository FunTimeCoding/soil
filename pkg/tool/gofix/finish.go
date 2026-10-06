package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func finish(
	w *workspace.Workspace,
	r *output.Results,
) {
	if w.Diff() {
		printChanges(w)

		return
	}

	changed, e := w.Commit()
	errors.PanicOnError(e)

	if len(changed) == 0 {
		return
	}

	for _, c := range r.Entries {
		c.Fixed = false
	}

	for _, path := range changed {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernConcurrentWrite,
				"changed since it was read - nothing written, run again",
				path,
				false,
			),
		)
	}
}
