package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddVirtualTag(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualTag,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Tag == "" {
		return response.Fail("virtual machine and tag are required")
	}

	result, e := s.service.AddVirtualTag(a.VirtualMachine, a.Tag)

	if e != nil {
		return s.answerMessage(x, e, "tag not added to virtual machine")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
