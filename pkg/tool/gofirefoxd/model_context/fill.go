package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Fill(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Fill,
) (*mcp.CallToolResult, error) {
	if a.Selector == "" {
		return response.Fail("selector is required")
	}

	v, e := s.client.Fill(
		int(a.TabIdentifier),
		a.Selector,
		a.Value,
		bool(a.Direct),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
