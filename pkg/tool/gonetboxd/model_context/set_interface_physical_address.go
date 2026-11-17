package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetInterfacePhysicalAddress(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SetInterfacePhysicalAddress,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Interface == "" || a.Address == "" {
		return response.Fail("device, interface and address are required")
	}

	result, e := s.service.SetInterfacePhysicalAddress(
		a.Device,
		a.Interface,
		a.Address,
	)

	if e != nil {
		return s.answerMessage(x, e, "physical address not set")
	}

	return response.SuccessAny(convert.InterfacePhysicalAddress(result))
}
