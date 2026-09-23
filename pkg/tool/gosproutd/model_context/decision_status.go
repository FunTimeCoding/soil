package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) decisionStatus(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DecisionStatus,
) (*mcp.CallToolResult, error) {
	open := map[string]int{}
	var pending []map[string]any

	for _, v := range s.service.Decisions(a.Session) {
		if v.State == constant.StateOpen && v.Resolution == "" {
			for _, f := range v.Frames {
				open[f.Name]++
			}

			continue
		}

		if v.ClearLine != "" {
			continue
		}

		pending = append(pending, pendingDecision(v))
	}

	return response.SuccessAny(
		map[string]any{"open_by_frame": open, "pending": pending},
	)
}
