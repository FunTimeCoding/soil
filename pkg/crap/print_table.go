package crap

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/crap/baseline"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/crap/report"
	"github.com/funtimecoding/soil/pkg/integers"
)

func printTable(
	r *report.Report,
	o *option.Report,
	b *baseline.Baseline,
) {
	t := newScoreTable(b)
	shown := 0

	for _, e := range r.Sorted() {
		if e.Score < o.Minimum || o.Top > 0 && shown >= o.Top {
			break
		}

		if b != nil && !o.All && e.Unchanged(o.Tolerance) {
			continue
		}

		shown++
		row := []string{
			mark(e.Score, o.Threshold),
			fmt.Sprintf(constant.ScoreFormat, e.Score),
			integers.ToString(e.Function.Complexity),
			fmt.Sprintf(constant.CoverageFormat, e.Coverage),
		}

		if b != nil {
			row = append(row, deltaColumn(e, o.Tolerance))
		}

		t.Add(append(row, label(e), e.Location(r.Root))...)
	}

	t.Print()
	console.Format(
		constant.SummaryFormat,
		len(r.Above(o.Threshold)),
		len(r.Entries),
		o.Threshold,
		r.Combined(),
		r.Average(),
	)

	if b != nil {
		console.Format(constant.BaselineFormat, r.Combined()-b.Combined)
	}

	if r.Skipped > 0 {
		console.Format(constant.SkippedFormat, r.Skipped)
	}
}
