package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreatePowerCable(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreatePowerCable,
) (*mcp.CallToolResult, error) {
	if a.OutletDevice == "" || a.Outlet == "" ||
		a.PortDevice == "" || a.Port == "" {
		return response.Fail(
			"outlet device, outlet, port device and port are required",
		)
	}

	result, e := s.service.CreatePowerCable(
		a.OutletDevice,
		a.Outlet,
		a.PortDevice,
		a.Port,
	)

	if e != nil {
		return s.answerMessage(x, e, "power cable not created")
	}

	return response.SuccessAny(convert.Cable(result))
}
