package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"

func (s *Service) ToolCalls(sessionIdentifier string) []tool_call.Call {
	return s.client.ToolCalls(sessionIdentifier)
}
