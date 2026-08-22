package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"time"
)

type sessionEntry struct {
	ModTime time.Time
	Size    int64
	Session *session.Session
}
