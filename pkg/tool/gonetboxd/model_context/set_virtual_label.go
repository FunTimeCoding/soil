package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetVirtualLabel(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualLabel,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Key == "" || a.Value == "" {
		return response.Fail("virtual machine, key and value are required")
	}

	result, e := s.service.SetVirtualLabel(a.VirtualMachine, a.Key, a.Value)

	if e != nil {
		return s.answerMessage(x, e, "label not set on virtual machine")
	}

	return response.SuccessAny(convert.Label(result))
}
