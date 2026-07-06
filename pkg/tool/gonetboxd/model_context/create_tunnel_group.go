package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTunnelGroup(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(parameter.Name)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	g, f := s.client.CreateTunnelGroup(name)

	if f != nil {
		return s.captureDetail(f)
	}

	return response.SuccessAny(convert.TunnelGroup(g))
}
