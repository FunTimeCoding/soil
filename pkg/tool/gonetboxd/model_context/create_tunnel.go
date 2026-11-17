package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTunnel(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateTunnel,
) (*mcp.CallToolResult, error) {
	if a.Name == "" || a.Encapsulation == "" || a.Group == "" {
		return response.Fail("name, encapsulation and group are required")
	}

	t, e := s.service.CreateTunnel(a.Name, a.Encapsulation, a.Group)

	if e != nil {
		return s.answerMessage(x, e, "tunnel not created")
	}

	return response.SuccessAny(convert.Tunnel(t))
}
