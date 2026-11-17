package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTunnelGroup(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Name,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	g, e := s.service.CreateTunnelGroup(a.Name)

	if e != nil {
		return s.answerMessage(x, e, "tunnel group not created")
	}

	return response.SuccessAny(convert.TunnelGroup(g))
}
