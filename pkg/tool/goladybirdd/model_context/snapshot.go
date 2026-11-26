package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Snapshot(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Snapshot,
) (*mcp.CallToolResult, error) {
	v, e := s.service.Snapshot(a.TabIdentifier)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.Success("%s", show(v))
}
