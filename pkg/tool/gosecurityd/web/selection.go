package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"net/http"
)

func selection(
	r *http.Request,
	fallback float64,
) (*system_selection.Selection, error) {
	m, e := minimum(r, fallback)

	if e != nil {
		return nil, e
	}

	q := r.URL.Query()
	result := system_selection.New().
		WithCoverage(q.Get(constant.ShowField)).
		WithMinimum(m).
		WithSearch(q.Get(constant.SearchField))

	for _, l := range q[constant.LabelField] {
		next, f := result.WithLabelText(l)

		if f != nil {
			return nil, f
		}

		result = next
	}

	return result, nil
}
