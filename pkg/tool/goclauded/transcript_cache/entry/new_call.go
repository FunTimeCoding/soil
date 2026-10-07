package entry

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"time"
)

func NewCall(
	modTime time.Time,
	size int64,
	calls []*tool_call.Call,
) *Call {
	return &Call{ModTime: modTime, Size: size, Calls: calls}
}
