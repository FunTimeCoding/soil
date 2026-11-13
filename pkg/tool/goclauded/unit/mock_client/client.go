package mock_client

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
)

type Client struct {
	sessions      map[string]*session.Session
	calls         map[string][]*tool_call.Call
	messages      map[string][]message.Message
	UserMessages  map[string][]string
	FirstMessages map[string]string
	Harbor        string
}
