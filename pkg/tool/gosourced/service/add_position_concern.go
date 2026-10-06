package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/token"
)

func addPositionConcern(
	r *output.Results,
	set *token.FileSet,
	position token.Pos,
	key string,
	fixed bool,
	format string,
	arguments ...any,
) {
	p := set.Position(position)
	r.AddConcern(
		concern.NewLine(
			key,
			fmt.Sprintf(format, arguments...),
			p.Filename,
			p.Line,
			"",
			fixed,
		),
	)
}
