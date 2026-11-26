package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/fill"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Fill(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Fill,
) (*mcp.CallToolResult, error) {
	r := fill.New(a.TabIdentifier, a.UID, a.Value)
	r.Direct = a.Direct != nil && *a.Direct
	r.Snapshot = a.Snapshot == nil || *a.Snapshot
	v, e := s.service.Fill(r)

	if e != nil {
		return s.captureDetail(x, e)
	}

	if !r.Snapshot {
		return response.Success("filled %s", a.UID)
	}

	return response.Success("filled %s\n\n%s", a.UID, show(v))
}
