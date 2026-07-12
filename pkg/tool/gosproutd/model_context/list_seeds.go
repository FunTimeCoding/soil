package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listSeeds(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	seeds := s.service.Seeds()
	var result []map[string]any

	for _, v := range seeds {
		result = append(
			result,
			map[string]any{
				"name":        v.Name,
				"path":        v.Path,
				"position":    v.Position,
				"modified_at": v.ModifiedAt,
			},
		)
	}

	return response.SuccessAny(result)
}
