package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_response"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/ollama/ollama/api"
)

func (c *Client) Chat(v *chat_request.Request) (*chat_response.Response, error) {
	r := v.Get()

	if r.Model == "" {
		r.Model = constant.Llama31
	}

	r.Stream = new(false)
	var result *api.ChatResponse
	e := c.client.Chat(
		c.context,
		r,
		func(r api.ChatResponse) error {
			result = &r

			return nil
		},
	)

	if e != nil {
		return nil, e
	}

	return chat_response.New(result), nil
}
