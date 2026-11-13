package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/layer"
)

func commitFixes(
	fixes *layer.Layer,
	o *option.Lint,
	r *output.Results,
) {
	fixes.BeforeCommit(o.BeforeCommit)
	refused, e := fixes.Commit(nil)
	errors.PanicOnError(e)

	if len(refused) == 0 {
		return
	}

	for _, c := range r.Entries {
		c.Fixed = false
	}

	for _, path := range refused {
		r.AddConcern(
			concern.NewFile(
				constant.ConcurrentWriteKey,
				constant.ConcurrentWriteText,
				path,
				false,
			),
		)
	}
}
