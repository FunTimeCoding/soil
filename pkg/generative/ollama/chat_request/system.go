package chat_request

import "github.com/funtimecoding/soil/pkg/generative/ollama/constant"

func (r *Request) System(s string) *Request {
	return r.Add(constant.SystemRole, s)
}
