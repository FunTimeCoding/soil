package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) setVirtualCustomField(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := r.RequireString(constant.VirtualMachine)

	if e != nil {
		return response.Fail("virtual_machine is required: %v", e)
	}

	field, f := r.RequireString(constant.Field)

	if f != nil {
		return response.Fail("field is required: %v", f)
	}

	result, g := s.service.SetVirtualMachineCustomFields(
		name,
		map[string]any{field: r.GetStringSlice(constant.Values, []string{})},
	)

	if not_found.Is(g) {
		return response.Fail("%v", g)
	}

	if g != nil {
		return s.captureFail(g, "custom field not set")
	}

	return response.SuccessAny(convert.VirtualMachine(result))
}
