package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListDeviceTags(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Device,
) (*mcp.CallToolResult, error) {
	if a.Device == "" {
		return response.Fail("device is required")
	}

	result, e := s.service.DeviceTagNames(a.Device)

	if e != nil {
		return s.answerMessage(x, e, "device tags not retrieved")
	}

	return response.SuccessAny(result)
}
