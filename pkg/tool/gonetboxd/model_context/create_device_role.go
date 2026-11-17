package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateDeviceRole(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateDeviceRole,
) (*mcp.CallToolResult, error) {
	if a.Name == "" {
		return response.Fail("name is required")
	}

	result, e := s.service.CreateDeviceRole(a.Name, a.VirtualMachineRole)

	if e != nil {
		return s.answerMessage(x, e, "device role not created")
	}

	return response.SuccessAny(convert.DeviceRole(result))
}
