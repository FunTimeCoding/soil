package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateVirtualDisk(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateVirtualDisk,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Name == "" {
		return response.Fail("virtual machine and name are required")
	}

	if a.Size <= 0 {
		return response.Fail("size must be positive")
	}

	result, e := s.service.CreateVirtualDisk(
		a.VirtualMachine,
		a.Name,
		int32(a.Size),
	)

	if e != nil {
		return s.answerMessage(x, e, "virtual disk not created")
	}

	return response.SuccessAny(convert.Disk(result))
}
