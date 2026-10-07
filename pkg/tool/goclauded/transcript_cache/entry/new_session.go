package entry

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"time"
)

func NewSession(
	modTime time.Time,
	size int64,
	session *session.Session,
) *Session {
	return &Session{ModTime: modTime, Size: size, Session: session}
}
