package crap

import (
	"github.com/funtimecoding/soil/pkg/crap/baseline"
	"github.com/funtimecoding/soil/pkg/crap/report"
)

func Compare(
	r *report.Report,
	b *baseline.Baseline,
) {
	for _, e := range r.Entries {
		if v, okay := b.Score(e.Function.Key()); okay {
			e.Previous = new(v)
		}
	}
}
