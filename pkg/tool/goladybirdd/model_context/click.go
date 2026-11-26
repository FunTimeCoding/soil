package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Click(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Click,
) (*mcp.CallToolResult, error) {
	observe := a.Snapshot == nil || *a.Snapshot
	v, e := s.service.Click(a.TabIdentifier, a.UID, observe)

	if e != nil {
		return s.captureDetail(x, e)
	}

	if !observe {
		return response.Success("%s", clicked(a.UID, v))
	}

	return response.Success("%s\n\n%s", clicked(a.UID, v), show(v))
}
