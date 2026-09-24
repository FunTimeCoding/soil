package crap

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/crap/baseline"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/mutation"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/crap/score"
	"github.com/funtimecoding/soil/pkg/notation"
)

func Run(o *option.Report) int {
	i := index.Load(o.Root, o.Patterns...)
	profile := o.Profile

	if profile == "" {
		profile = coverage.Run(o.Root, "", o.Patterns...)
	}

	r := Build(
		i,
		coverage.Functions(o.Root, profile),
		score.ParsePolicy(o.Missing),
	)

	if o.Mutation != "" {
		Annotate(r, mutation.Load(o.Mutation))
	}

	var b *baseline.Baseline

	if o.Baseline != "" {
		b = baseline.Load(o.Baseline)
		Compare(r, b)
	}

	if o.Notation {
		console.Line(notation.MarshalIndent(r))
	} else {
		printTable(r, o, b)
	}

	if o.FailAbove && len(r.Above(o.Threshold)) > 0 {
		console.Format(
			constant.FailFormat,
			len(r.Above(o.Threshold)),
			o.Threshold,
		)

		return 1
	}

	if o.FailRegression && b != nil {
		if v := r.Regressions(o.Tolerance, o.IgnoreCovered); len(v) > 0 {
			printRegressions(r, v)
			console.Format(constant.RegressionFormat, len(v))

			return 1
		}
	}

	return 0
}
