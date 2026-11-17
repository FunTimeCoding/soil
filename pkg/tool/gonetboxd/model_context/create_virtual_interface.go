package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateVirtualInterface(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateVirtualInterface,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Name == "" {
		return response.Fail("virtual machine and name are required")
	}

	result, e := s.service.CreateVirtualInterface(a.VirtualMachine, a.Name)

	if e != nil {
		return s.answerMessage(x, e, "virtual interface not created")
	}

	return response.SuccessAny(convert.VirtualInterface(result))
}
