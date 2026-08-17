package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"time"
)

func callTime(c *tool_call.Call) time.Time {
	result, e := time.Parse(time.RFC3339, c.Timestamp)

	if e != nil {
		return time.Time{}
	}

	return result
}
