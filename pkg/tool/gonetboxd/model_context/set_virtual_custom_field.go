package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetVirtualCustomField(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualCustomField,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Field == "" {
		return response.Fail("virtual machine and field are required")
	}

	values := a.Values

	if values == nil {
		values = []string{}
	}

	result, e := s.service.SetVirtualMachineCustomFields(
		a.VirtualMachine,
		map[string]any{a.Field: values},
	)

	if e != nil {
		return s.answerMessage(x, e, "custom field not set")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
