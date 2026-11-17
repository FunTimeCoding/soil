package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetDeviceLabel(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceLabel,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Key == "" || a.Value == "" {
		return response.Fail("device, key and value are required")
	}

	result, e := s.service.SetDeviceLabel(a.Device, a.Key, a.Value)

	if e != nil {
		return s.answerMessage(x, e, "label not set on device")
	}

	return response.SuccessAny(convert.Label(result))
}
