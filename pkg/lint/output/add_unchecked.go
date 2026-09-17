package output

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (r *Results) AddUnchecked(
	path string,
	line int,
	span string,
	reason constant.Reason,
) {
	r.Unchecked = append(
		r.Unchecked,
		&Unchecked{
			Path:   r.Relativize(path),
			Line:   line,
			Span:   span,
			Reason: reason,
		},
	)
}
