package unchecked

import "github.com/funtimecoding/soil/pkg/lint/constant"

func New(
	path string,
	line int,
	span string,
	reason constant.Reason,
) *Unchecked {
	return &Unchecked{Path: path, Line: line, Span: span, Reason: reason}
}
