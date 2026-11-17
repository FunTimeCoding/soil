package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddDeviceTag(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceTag,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Tag == "" {
		return response.Fail("device and tag are required")
	}

	result, e := s.service.AddTag(a.Device, a.Tag)

	if e != nil {
		return s.answerMessage(x, e, "tag not added to device")
	}

	return response.SuccessAny(convert.Device(result))
}
