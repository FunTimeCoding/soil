package mock_client

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_response"
	"github.com/ollama/ollama/api"
)

func (c *Client) Chat(
	v *chat_request.Request,
) (*chat_response.Response, error) {
	c.Chats = append(c.Chats, v)
	var last string

	if m := v.Get().Messages; len(m) > 0 {
		last = m[len(m)-1].Content
	}

	text := c.answer(last)

	if c.Failure != nil {
		return nil, c.Failure
	}

	return chat_response.New(
		&api.ChatResponse{
			Message: api.Message{Role: constant.RoleAssistant, Content: text},
		},
	), nil
}
