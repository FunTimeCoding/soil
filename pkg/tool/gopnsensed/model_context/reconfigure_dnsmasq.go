package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) reconfigureDnsmasq(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	if e := s.service.ReconfigureDnsmasq(); e != nil {
		return s.answer(x, e)
	}

	return response.Success(constant.ReconfigureDnsmasq)
}
