package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateCable(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateCable,
) (*mcp.CallToolResult, error) {
	if a.DeviceA == "" || a.InterfaceA == "" ||
		a.DeviceB == "" || a.InterfaceB == "" {
		return response.Fail("both devices and both interfaces are required")
	}

	result, e := s.service.CreateCable(
		a.DeviceA,
		a.InterfaceA,
		a.DeviceB,
		a.InterfaceB,
	)

	if e != nil {
		return s.answerMessage(x, e, "cable not created")
	}

	return response.SuccessAny(convert.Cable(result))
}
