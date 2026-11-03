package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
)

func reportAliasRename(
	d *import_decision.Decision,
	filePath string,
	r *output.Results,
) {
	e := d.Entry
	done := "re-aliased"

	if d.Want == e.Declared {
		done = "de-aliased"
	}

	r.AddConcern(
		concern.NewFile(
			"import_alias",
			fmt.Sprintf("%s %s → %s (%s)", done, e.Local, d.Want, e.Path),
			filePath,
			true,
		),
	)
}
