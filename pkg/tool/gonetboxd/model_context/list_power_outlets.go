package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListPowerOutlets(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Device,
) (*mcp.CallToolResult, error) {
	if a.Device == "" {
		return response.Fail("device is required")
	}

	result, e := s.service.DevicePowerOutlets(a.Device)

	if e != nil {
		return s.answerMessage(x, e, "power outlets not retrieved")
	}

	return response.SuccessAny(convert.PowerOutlets(result))
}
