package chat_request

import "github.com/ollama/ollama/api"

func (r *Request) Add(
	role string,
	content string,
) *Request {
	r.request.Messages = append(
		r.request.Messages,
		api.Message{Role: role, Content: content},
	)

	return r
}
