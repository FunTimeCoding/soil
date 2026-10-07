package session_tool_count

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func New(
	session *session.Session,
	count int,
) *Count {
	return &Count{Session: session, Count: count}
}
