package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_response"
)

func (c *Client) ChatSimple(p string) *chat_response.Response {
	return c.MustChat(chat_request.New().User(p))
}
