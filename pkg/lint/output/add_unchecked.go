package output

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

func (r *Results) AddUnchecked(
	path string,
	line int,
	span string,
	reason constant.Reason,
) {
	r.Unchecked = append(
		r.Unchecked,
		unchecked.New(r.Relativize(path), line, span, reason))
}
