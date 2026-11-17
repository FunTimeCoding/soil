package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTunnelGroups(
	x context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.service.TunnelGroups()

	if e != nil {
		return s.answerMessage(x, e, "tunnel groups not listed")
	}

	return response.SuccessAny(convert.TunnelGroups(result))
}
