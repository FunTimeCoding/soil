package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"

func (c *Client) AddToolCall(
	sessionIdentifier string,
	call *tool_call.Call,
) {
	c.calls[sessionIdentifier] = append(c.calls[sessionIdentifier], call)
}
