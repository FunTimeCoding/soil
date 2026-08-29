package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) reconfigureDnsmasq(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	if e := s.opnsense.ReconfigureDnsmasq(); e != nil {
		return s.captureDetail(e)
	}

	return response.Success(constant.ReconfigureDnsmasq)
}
