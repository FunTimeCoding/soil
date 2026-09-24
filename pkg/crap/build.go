package crap

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/entry"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/report"
)

func Build(
	i *index.Index,
	coverage map[string]float64,
	p constant.Policy,
) *report.Report {
	result := report.New(i.Root)

	for _, f := range i.Functions {
		c, okay := coverage[f.Key()]

		switch {
		case okay:
			result.Add(entry.New(f, c))
		case p == constant.Skip:
			result.Skipped++
		default:
			result.Add(entry.NewMissing(f, p))
		}
	}

	return result
}
