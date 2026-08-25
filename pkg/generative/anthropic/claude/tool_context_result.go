package claude

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"

type ToolContextResult struct {
	ToolName       string
	ToolIdentifier string
	Before         []message.Message
	After          []message.Message
}
