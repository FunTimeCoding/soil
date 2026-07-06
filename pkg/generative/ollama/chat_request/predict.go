package chat_request

import "github.com/funtimecoding/soil/pkg/generative/ollama/constant"

func (r *Request) Predict(size int) *Request {
	return r.Option(constant.PredictSize, size)
}
