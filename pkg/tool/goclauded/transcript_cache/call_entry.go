package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"time"
)

type callEntry struct {
	ModTime time.Time
	Size    int64
	Calls   []tool_call.Call
}
