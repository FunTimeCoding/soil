package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetInterfaceEthernetPower(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SetInterfaceEthernetPower,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Interface == "" {
		return response.Fail("device and interface are required")
	}

	result, e := s.service.SetInterfaceEthernetPower(
		a.Device,
		a.Interface,
		a.Mode,
		a.Type,
	)

	if e != nil {
		return s.answerMessage(x, e, "ethernet power not set")
	}

	return response.SuccessAny(convert.Interface(result))
}
