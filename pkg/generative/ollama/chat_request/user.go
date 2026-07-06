package chat_request

import "github.com/funtimecoding/soil/pkg/generative/ollama/constant"

func (r *Request) User(s string) *Request {
	return r.Add(constant.UserRole, s)
}
