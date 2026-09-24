package result

import (
	"github.com/funtimecoding/soil/pkg/measure/summary"
	"sort"
)

func (r *Result) ByLanguage() []*summary.Summary {
	index := map[string]*summary.Summary{}
	var out []*summary.Summary

	for _, f := range r.Files {
		s, okay := index[f.Language]

		if !okay {
			s = summary.New(f.Language)
			index[f.Language] = s
			out = append(out, s)
		}

		s.Add(f.Count)
	}

	sort.SliceStable(
		out,
		func(
			i int,
			j int,
		) bool {
			if out[i].Count.Code == out[j].Count.Code {
				return out[i].Language < out[j].Language
			}

			return out[i].Count.Code > out[j].Count.Code
		},
	)

	return out
}
