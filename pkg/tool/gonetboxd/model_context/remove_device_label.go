package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) RemoveDeviceLabel(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceLabelKey,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Key == "" {
		return response.Fail("device and key are required")
	}

	if e := s.service.RemoveDeviceLabel(a.Device, a.Key); e != nil {
		return s.answerMessage(x, e, "label not removed from device")
	}

	return response.Success("label removed")
}
