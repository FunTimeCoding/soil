package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) pushDecision(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.PushDecision,
) (*mcp.CallToolResult, error) {
	d, e := s.service.PushDecision(
		a.Session,
		a.Question,
		a.DefaultAction,
		a.Options,
		a.Frames,
	)

	if e != nil {
		return response.Fail("%s", e)
	}

	return response.SuccessAny(
		map[string]any{"identifier": d.Identifier, "state": d.State},
	)
}
