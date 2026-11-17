package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetVirtualMachine(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualMachine,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" {
		return response.Fail("virtual machine is required")
	}

	m, labels, e := s.service.VirtualMachine(a.VirtualMachine)

	if e != nil {
		return s.answerMessage(x, e, "virtual machine not retrieved")
	}

	return response.SuccessAny(convert.LabeledVirtualMachine(m, labels))
}
