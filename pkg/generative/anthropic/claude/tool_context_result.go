package claude

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"

type ToolContextResult struct {
	ToolName string
	ToolID   string
	Before   []message.Message
	After    []message.Message
}
