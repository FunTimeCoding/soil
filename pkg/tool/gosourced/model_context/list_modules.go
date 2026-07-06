package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listModules(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	modules := s.service.Modules()
	var result []map[string]any

	for _, m := range modules {
		result = append(
			result,
			map[string]any{
				"name":      m.Name,
				"directory": m.Directory,
			},
		)
	}

	return response.SuccessAny(result)
}
