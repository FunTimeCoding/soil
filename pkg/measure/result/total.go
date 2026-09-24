package result

import (
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/summary"
)

func (r *Result) Total() *summary.Summary {
	result := summary.New(constant.RowTotal)

	for _, f := range r.Files {
		result.Add(f.Count)
	}

	return result
}
