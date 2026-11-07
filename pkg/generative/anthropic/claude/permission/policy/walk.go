package policy

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"

type Walk struct {
	rules     []*allowlist.Rule
	roots     []string
	origin    string
	base      string
	moved     bool
	variables map[string]string
	commands  int
	classes   []string
	sources   []string
	blockers  []string
}
