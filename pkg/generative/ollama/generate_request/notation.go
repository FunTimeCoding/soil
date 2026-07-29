package generate_request

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (r *Request) Notation() *Request {
	r.request.Format = notation.Marshal(constant.OllamaNotationFormat)

	return r
}
