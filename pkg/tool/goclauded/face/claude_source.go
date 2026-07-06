package face

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/peek"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
)

type ClaudeSource interface {
	Sessions() []*session.Session
	Resolve(query string) *session.Session
	Messages(sessionIdentifier string) []message.Message
	FirstUserMessage(sessionIdentifier string) string
	Peek(sessionIdentifier string) *peek.Peek
	Delete(sessionIdentifier string)
	SessionsByTool(toolFilter string) []*claude.SessionToolCount
	ToolCalls(sessionIdentifier string) []tool_call.Call
	ToolContext(
		sessionIdentifier string,
		toolFilter string,
		surroundCount int,
	) []claude.ToolContextResult
}
