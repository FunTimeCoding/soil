package tracker

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *State) recordCall(
	b *notation.ContentBlock,
	timestamp string,
) {
	if s.Pending == nil {
		s.Pending = map[string]tool_call.Call{}
	}

	if len(s.Pending) >= constant.ClaudePendingCallLimit {
		s.dropOldestPending()
	}

	s.Pending[b.Identifier] = tool_call.Call{
		Name:       b.Name,
		Identifier: b.Identifier,
		Timestamp:  timestamp,
		Input:      string(b.Input),
	}
}
