package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListDeviceTypes(
	x context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	result, e := s.service.DeviceTypes()

	if e != nil {
		return s.answerMessage(x, e, "device types not listed")
	}

	return response.SuccessAny(convert.DeviceTypes(result))
}
