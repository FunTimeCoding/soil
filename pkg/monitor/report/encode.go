package report

import "github.com/funtimecoding/soil/pkg/notation"

func (r *Report) Encode() string {
	r.Count = len(r.Items)

	return notation.Encode(r, true)
}
