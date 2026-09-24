package report

import (
	"github.com/funtimecoding/soil/pkg/crap/entry"
	"sort"
)

func (r *Report) Sorted() []*entry.Entry {
	out := make([]*entry.Entry, len(r.Entries))
	copy(out, r.Entries)
	sort.SliceStable(
		out,
		func(
			i int,
			j int,
		) bool {
			if out[i].Score != out[j].Score {
				return out[i].Score > out[j].Score
			}

			return out[i].Function.Key() < out[j].Function.Key()
		},
	)

	return out
}
