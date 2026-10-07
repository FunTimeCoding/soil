package service

import "github.com/funtimecoding/soil/pkg/generative/types/session_tool_count"

func (s *Service) SessionsByTool(toolFilter string) []*session_tool_count.Count {
	return s.client.SessionsByTool(toolFilter)
}
