package service

import "github.com/funtimecoding/soil/pkg/generative/types/tool_context_result"

func (s *Service) ToolContext(
	sessionIdentifier string,
	toolFilter string,
	surroundCount int,
) []tool_context_result.Result {
	return s.client.ToolContext(sessionIdentifier, toolFilter, surroundCount)
}
