package tracker

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"

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
}
