package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) cruiseStatus(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CruiseStatus,
) (*mcp.CallToolResult, error) {
	setting := s.service.CruiseSetting(a.Session)
	result := map[string]any{"mode": setting.Mode}

	if setting.Pace > 0 {
		result["pace_minutes"] = setting.Pace.Minutes()
	}

	if setting.PromotedAt != nil {
		result["last_promoted_at"] = setting.PromotedAt
	}

	return response.SuccessAny(result)
}
