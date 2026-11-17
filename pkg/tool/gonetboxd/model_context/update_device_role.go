package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateDeviceRole(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateDeviceRole,
) (*mcp.CallToolResult, error) {
	if a.Role == "" {
		return response.Fail("role is required")
	}

	result, e := s.service.UpdateDeviceRole(a.Role, filled(a.Name))

	if e != nil {
		return s.answerMessage(x, e, "device role not updated")
	}

	return response.SuccessAny(convert.DeviceRole(result))
}
