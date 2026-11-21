package web

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"net/http"
	"strings"
)

func selection(r *http.Request) (*system_selection.Selection, error) {
	m, e := minimum(r, 0)

	if e != nil {
		return nil, e
	}

	q := r.URL.Query()
	result := system_selection.New()
	result.Show = q.Get(constant.ShowField)
	result.Minimum = m
	result.Search = q.Get(constant.SearchField)

	for _, l := range q[constant.LabelField] {
		key, value, _ := strings.Cut(l, constant.LabelSeparator)

		if key == "" || value == "" {
			return nil, validation.New("label filter malformed: %s", l)
		}

		result.Labels[key] = value
	}

	return result, nil
}
