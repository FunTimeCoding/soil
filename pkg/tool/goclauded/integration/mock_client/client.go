package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

type Client struct {
	sessions      map[string]*session.Session
	UserMessages  map[string][]string
	FirstMessages map[string]string
}
