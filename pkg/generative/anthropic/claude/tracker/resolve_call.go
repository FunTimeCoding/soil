package tracker

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
)

func (s *State) resolveCall(b *notation.ContentBlock) *tool_call.Call {
	c, found := s.Pending[b.ToolUseIdentifier]

	if !found {
		return nil
	}

	delete(s.Pending, b.ToolUseIdentifier)
	c.Result = string(b.Content)

	return &c
}
