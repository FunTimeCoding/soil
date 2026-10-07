package session_tool_count

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

type Count struct {
	Session *session.Session
	Count   int
}
