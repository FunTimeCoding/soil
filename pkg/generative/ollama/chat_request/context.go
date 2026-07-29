package chat_request

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (r *Request) Context(size int) *Request {
	return r.Option(constant.OllamaContextSize, size)
}
