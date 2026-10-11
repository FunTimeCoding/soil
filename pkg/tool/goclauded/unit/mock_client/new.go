package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func New() *Client {
	return &Client{
		sessions:      make(map[string]*session.Session),
		UserMessages:  make(map[string][]string),
		FirstMessages: make(map[string]string),
	}
}
