package tracker

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
)

type State struct {
	Offset           int64
	Lines            int
	FirstTimestamp   string
	LastTimestamp    string
	Slug             string
	WorkDirectory    string
	Branch           string
	UserMessageCount int
	FirstMessage     string
	Usage            map[string]*pricing.Tokens
	RecentMessages   []string
	Pending          map[string]tool_call.Call
}
