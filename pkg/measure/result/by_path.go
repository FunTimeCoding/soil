package result

import (
	"github.com/funtimecoding/soil/pkg/measure/file"
	"sort"
)

func (r *Result) ByPath() []*file.File {
	out := make([]*file.File, len(r.Files))
	copy(out, r.Files)
	sort.SliceStable(
		out,
		func(
			i int,
			j int,
		) bool {
			return out[i].Path < out[j].Path
		},
	)

	return out
}
