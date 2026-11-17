package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateVirtualAddress(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateVirtualAddress,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Interface == "" || a.Address == "" {
		return response.Fail(
			"virtual machine, interface and address are required",
		)
	}

	result, e := s.service.CreateVirtualAddress(
		a.VirtualMachine,
		a.Interface,
		a.Address,
		a.Status,
	)

	if e != nil {
		return s.answerMessage(x, e, "address not assigned")
	}

	return response.SuccessAny(convert.Address(result))
}
