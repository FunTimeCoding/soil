package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func addSuggestion(
	r *output.Results,
	filePath string,
	text string,
) {
	c := concern.NewFile("import_alias", text, filePath, false)
	c.Planned = true
	r.AddConcern(c)
}
