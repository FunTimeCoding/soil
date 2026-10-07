package tool_context_result

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"

type Result struct {
	ToolName       string
	ToolIdentifier string
	Before         []message.Message
	After          []message.Message
}
