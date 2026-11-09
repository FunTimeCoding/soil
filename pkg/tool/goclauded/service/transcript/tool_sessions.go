package transcript

import "github.com/funtimecoding/soil/pkg/generative/types/session_tool_count"

func (s *Service) ToolSessions(toolFilter string) []*session_tool_count.Count {
	return s.client.ToolSessions(toolFilter)
}
