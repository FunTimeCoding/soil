package output

import "github.com/funtimecoding/soil/pkg/lint/constant"

type Unchecked struct {
	Path   string
	Line   int
	Span   string
	Reason constant.Reason
}
