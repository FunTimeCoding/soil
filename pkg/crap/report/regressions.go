package report

import "github.com/funtimecoding/soil/pkg/crap/entry"

func (r *Report) Regressions(
	tolerance float64,
	ignoreCovered bool,
) []*entry.Entry {
	var out []*entry.Entry

	for _, e := range r.Entries {
		if !e.Regressed(tolerance) {
			continue
		}

		if ignoreCovered && e.Coverage >= 100 {
			continue
		}

		out = append(out, e)
	}

	return out
}
