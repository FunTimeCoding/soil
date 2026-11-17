package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreatePowerPort(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceComponent,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Name == "" {
		return response.Fail("device and name are required")
	}

	result, e := s.service.CreatePowerPort(a.Device, a.Name)

	if e != nil {
		return s.answerMessage(x, e, "power port not created")
	}

	return response.SuccessAny(convert.PowerPort(result))
}
