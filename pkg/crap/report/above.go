package report

import "github.com/funtimecoding/soil/pkg/crap/entry"

func (r *Report) Above(threshold float64) []*entry.Entry {
	var out []*entry.Entry

	for _, e := range r.Entries {
		if e.Score > threshold {
			out = append(out, e)
		}
	}

	return out
}
