package crap

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/entry"
	"github.com/funtimecoding/soil/pkg/crap/report"
)

func printRegressions(
	r *report.Report,
	v []*entry.Entry,
) {
	console.Line(constant.RegressionHeader)
	t := table.New(
		constant.ColumnLocation,
		constant.ColumnFunction,
		constant.ColumnPrevious,
		constant.ColumnScore,
		constant.ColumnDelta,
	).Right(2, 3, 4)

	for _, e := range v {
		t.Add(
			e.Location(r.Root),
			e.Function.QualifiedName(),
			fmt.Sprintf(constant.ScoreFormat, *e.Previous),
			fmt.Sprintf(constant.ScoreFormat, e.Score),
			fmt.Sprintf(constant.DeltaFormat, e.Delta()),
		)
	}

	t.Print()
}
