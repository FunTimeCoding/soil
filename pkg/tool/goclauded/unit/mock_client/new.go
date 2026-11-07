package mock_client

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
)

func New() *Client {
	return &Client{
		sessions:      make(map[string]*session.Session),
		calls:         make(map[string][]*tool_call.Call),
		UserMessages:  make(map[string][]string),
		FirstMessages: make(map[string]string),
	}
}
