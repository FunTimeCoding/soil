package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func reportAlias(
	r *output.Results,
	filePath string,
	suggestion string,
	applied string,
	realias bool,
) {
	if realias {
		r.AddConcern(concern.NewFile("import_alias", applied, filePath, true))

		return
	}

	addSuggestion(r, filePath, suggestion)
}
