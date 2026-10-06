package request

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (q *Request) WithNotation(v any) *Request {
	return q.WithBody(constant.Object, notation.Marshal(v))
}
