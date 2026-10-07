package entry

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"time"
)

type Session struct {
	ModTime time.Time
	Size    int64
	Session *session.Session
}
