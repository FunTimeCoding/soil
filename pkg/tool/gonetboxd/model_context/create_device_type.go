package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateDeviceType(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateDeviceType,
) (*mcp.CallToolResult, error) {
	if a.Model == "" || a.Manufacturer == "" {
		return response.Fail("model and manufacturer are required")
	}

	result, e := s.service.CreateDeviceType(a.Model, a.Manufacturer)

	if e != nil {
		return s.answerMessage(x, e, "device type not created")
	}

	return response.SuccessAny(convert.DeviceType(result))
}
