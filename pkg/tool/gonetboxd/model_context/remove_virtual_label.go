package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) RemoveVirtualLabel(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualLabelKey,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Key == "" {
		return response.Fail("virtual machine and key are required")
	}

	if e := s.service.RemoveVirtualLabel(a.VirtualMachine, a.Key); e != nil {
		return s.answerMessage(x, e, "label not removed from virtual machine")
	}

	return response.Success("label removed")
}
