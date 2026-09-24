package baseline

import "github.com/funtimecoding/soil/pkg/crap/report"

func New(r *report.Report) *Baseline {
	result := &Baseline{scores: map[string]float64{}}

	for _, e := range r.Entries {
		result.scores[e.Function.Key()] = e.Score
		result.Combined += e.Score
	}

	return result
}
